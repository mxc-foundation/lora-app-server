package gwping

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/mxc-foundation/lpwan-app-server/internal/backend/networkserver"
	"github.com/mxc-foundation/lpwan-app-server/internal/logging"
	"github.com/mxc-foundation/lpwan-app-server/internal/storage"
	"github.com/mxc-foundation/lpwan-server/api/as"
	"github.com/mxc-foundation/lpwan-server/api/ns"
	"github.com/brocaar/lorawan"
)

const (
	micLookupExpire = time.Second * 10
	micLookupTempl  = "lora:as:gwping:%s"
)

// SendPingLoop is a never returning function sending the gateway pings.
func SendPingLoop() {
	for {
		ctxID, err := uuid.NewV4()
		if err != nil {
			log.WithError(err).Error("new uuid error")
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, logging.ContextIDKey, ctxID)

		if err := sendGatewayPing(ctx); err != nil {
			log.Errorf("send gateway ping error: %s", err)
		}
		time.Sleep(time.Second)
	}
}

// HandleReceivedPing handles a ping received by one or multiple gateways.
func HandleReceivedPing(ctx context.Context, req *as.HandleProprietaryUplinkRequest) error {
	var mic lorawan.MIC
	copy(mic[:], req.Mic)

	id, err := getPingLookup(mic)
	if err != nil {
		return errors.Wrap(err, "get ping lookup error")
	}

	if err = deletePingLookup(mic); err != nil {
		log.Errorf("delete ping lookup error: %s", err)
	}

	ping, err := storage.GetGatewayPing(ctx, storage.DB(), id)
	if err != nil {
		return errors.Wrap(err, "get gateway ping error")
	}

	err = storage.Transaction(func(tx sqlx.Ext) error {
		for _, rx := range req.RxInfo {
			var mac lorawan.EUI64
			copy(mac[:], rx.GatewayId)

			// ignore pings received by the sending gateway
			if ping.GatewayMAC == mac {
				continue
			}

			var receivedAt *time.Time
			if rx.Time != nil {
				ts, err := ptypes.Timestamp(rx.Time)
				if err != nil {
					return err
				}
				receivedAt = &ts
			}

			pingRX := storage.GatewayPingRX{
				PingID:     id,
				GatewayMAC: mac,
				ReceivedAt: receivedAt,
				RSSI:       int(rx.Rssi),
				LoRaSNR:    rx.LoraSnr,
			}

			if rx.Location != nil {
				pingRX.Location = storage.GPSPoint{
					Latitude:  rx.Location.Latitude,
					Longitude: rx.Location.Longitude,
				}
				pingRX.Altitude = rx.Location.Altitude
			}

			err := storage.CreateGatewayPingRX(ctx, tx, &pingRX)
			if err != nil {
				return errors.Wrap(err, "create gateway ping rx error")
			}
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "transaction error")
	}

	return nil
}

// sendGatewayPing selects the next gateway to ping, creates the "ping"
// frame and sends this frame to the network-server for transmission.
func sendGatewayPing(ctx context.Context) error {
	return storage.Transaction(func(tx sqlx.Ext) error {
		gw, err := getGatewayForPing(tx)
		if err != nil {
			return errors.Wrap(err, "get gateway for ping error")
		}
		if gw == nil {
			return nil
		}

		n, err := storage.GetNetworkServer(ctx, storage.DB(), gw.NetworkServerID)
		if err != nil {
			return errors.Wrap(err, "get network-server error")
		}

		ping := storage.GatewayPing{
			GatewayMAC: gw.MAC,
			Frequency:  n.GatewayDiscoveryTXFrequency,
			DR:         n.GatewayDiscoveryDR,
		}
		err = storage.CreateGatewayPing(ctx, tx, &ping)
		if err != nil {
			return errors.Wrap(err, "create gateway ping error")
		}

		var mic lorawan.MIC
		if _, err = rand.Read(mic[:]); err != nil {
			return errors.Wrap(err, "read random bytes error")
		}

		err = CreatePingLookup(mic, ping.ID)
		if err != nil {
			return errors.Wrap(err, "store mic lookup error")
		}

		err = sendPing(mic, n, ping)
		if err != nil {
			return errors.Wrap(err, "send ping error")
		}

		gw.LastPingID = &ping.ID
		gw.LastPingSentAt = &ping.CreatedAt

		err = storage.UpdateGateway(ctx, tx, gw)
		if err != nil {
			return errors.Wrap(err, "update gateway error")
		}

		return nil
	})
}

// getGatewayForPing returns the next gateway for sending a ping. If no gateway
// matches the filter criteria, nil is returned.
func getGatewayForPing(tx sqlx.Ext) (*storage.Gateway, error) {
	var gw storage.Gateway

	err := sqlx.Get(tx, &gw, `
		select
			g.*
		from gateway g
		inner join network_server ns
			on ns.id = g.network_server_id
		where
			ns.gateway_discovery_enabled = true
			and g.ping = true
			and (g.last_ping_sent_at is null or g.last_ping_sent_at <= (now() - (interval '24 hours' / ns.gateway_discovery_interval)))
		order by last_ping_sent_at
		limit 1
		for update`,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "select error")
	}

	return &gw, nil
}

func sendPing(mic lorawan.MIC, n storage.NetworkServer, ping storage.GatewayPing) error {
	nsClient, err := networkserver.GetPool().Get(n.Server, []byte(n.CACert), []byte(n.TLSCert), []byte(n.TLSKey))
	if err != nil {
		return errors.Wrap(err, "get network-server client error")
	}

	_, err = nsClient.SendProprietaryPayload(context.Background(), &ns.SendProprietaryPayloadRequest{
		Mic:                   mic[:],
		GatewayMacs:           [][]byte{ping.GatewayMAC[:]},
		PolarizationInversion: false,
		Frequency:             uint32(ping.Frequency),
		Dr:                    uint32(ping.DR),
	})
	if err != nil {
		return errors.Wrap(err, "send proprietary payload error")
	}

	log.WithFields(log.Fields{
		"gateway_mac": ping.GatewayMAC,
		"id":          ping.ID,
	}).Info("gateway ping sent to network-server")

	return nil
}

// CreatePingLookup creates an automatically expiring MIC to ping id lookup.
func CreatePingLookup(mic lorawan.MIC, id int64) error {
	c := storage.RedisPool().Get()
	defer c.Close()

	_, err := redis.String(c.Do("PSETEX", fmt.Sprintf(micLookupTempl, mic), int64(micLookupExpire)/int64(time.Millisecond), id))
	if err != nil {
		return errors.Wrap(err, "set mic lookup error")
	}
	return nil
}

func getPingLookup(mic lorawan.MIC) (int64, error) {
	c := storage.RedisPool().Get()
	defer c.Close()

	id, err := redis.Int64(c.Do("GET", fmt.Sprintf(micLookupTempl, mic)))
	if err != nil {
		return 0, errors.Wrap(err, "get ping lookup error")
	}

	return id, nil
}

func deletePingLookup(mic lorawan.MIC) error {
	c := storage.RedisPool().Get()
	defer c.Close()

	_, err := redis.Int(c.Do("DEL", fmt.Sprintf(micLookupTempl, mic)))
	if err != nil {
		return errors.Wrap(err, "delete ping lookup error")
	}

	return nil
}
