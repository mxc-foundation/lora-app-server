package uplink

import (
	"context"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"time"

	keywrap "github.com/NickBall/go-aes-key-wrap"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/brocaar/chirpstack-api/go/v3/as"
	pb "github.com/brocaar/chirpstack-api/go/v3/as/integration"
	"github.com/brocaar/chirpstack-api/go/v3/common"
	"github.com/brocaar/lorawan"
	"github.com/brocaar/lorawan/gps"

	"github.com/mxc-foundation/lpwan-app-server/internal/applayer/clocksync"
	"github.com/mxc-foundation/lpwan-app-server/internal/applayer/fragmentation"
	"github.com/mxc-foundation/lpwan-app-server/internal/applayer/multicastsetup"
	"github.com/mxc-foundation/lpwan-app-server/internal/codec"
	"github.com/mxc-foundation/lpwan-app-server/internal/integration"
	joinserver "github.com/mxc-foundation/lpwan-app-server/internal/js"
	"github.com/mxc-foundation/lpwan-app-server/internal/logging"

	"github.com/mxc-foundation/lpwan-app-server/internal/integration/models"
	apps "github.com/mxc-foundation/lpwan-app-server/internal/modules/application/data"
	dps "github.com/mxc-foundation/lpwan-app-server/internal/modules/device-profile/data"
	ds "github.com/mxc-foundation/lpwan-app-server/internal/modules/device/data"
	"github.com/mxc-foundation/lpwan-app-server/internal/storage/store"
)

type uplinkContext struct {
	handler *store.Handler

	uplinkDataReq as.HandleUplinkDataRequest

	ctx           context.Context
	device        ds.Device
	application   apps.Application
	deviceProfile dps.DeviceProfile

	data          []byte
	objectJSON    string
	gIntegrations []models.IntegrationHandler
}

var tasks = []func(*uplinkContext) error{
	getDevice,
	getApplication,
	getDeviceProfile,
	updateDeviceLastSeenAndDR,
	updateDeviceActivation,
	decryptPayload,
	handleApplicationLayers,
	handleCodec,
	handleIntegrations,
}

// Handle handles the uplink event.
func Handle(ctx context.Context, req as.HandleUplinkDataRequest, handler *store.Handler,
	gIntegrations []models.IntegrationHandler) error {
	uc := uplinkContext{
		handler:       handler,
		ctx:           ctx,
		uplinkDataReq: req,
		gIntegrations: gIntegrations,
	}

	for _, f := range tasks {
		if err := f(&uc); err != nil {
			if err == ErrAbort {
				return nil
			}
			return err
		}
	}

	return nil
}

func getDevice(ctx *uplinkContext) error {
	var err error
	var devEUI lorawan.EUI64
	copy(devEUI[:], ctx.uplinkDataReq.DevEui)

	ctx.device, err = ctx.handler.GetDevice(ctx.ctx, devEUI, false)
	if err != nil {
		return errors.Wrap(err, "get device error")
	}
	return nil
}

func getDeviceProfile(ctx *uplinkContext) error {
	var err error
	ctx.deviceProfile, err = ctx.handler.GetDeviceProfile(ctx.ctx, ctx.device.DeviceProfileID, false)
	if err != nil {
		return errors.Wrap(err, "get device-profile error")
	}
	return nil
}

func getApplication(ctx *uplinkContext) error {
	var err error
	ctx.application, err = ctx.handler.GetApplication(ctx.ctx, ctx.device.ApplicationID)
	if err != nil {
		return errors.Wrap(err, "get application error")
	}
	return nil
}

func updateDeviceLastSeenAndDR(ctx *uplinkContext) error {
	if err := ctx.handler.UpdateDeviceLastSeenAndDR(ctx.ctx, ctx.device.DevEUI, time.Now(), int(ctx.uplinkDataReq.Dr)); err != nil {
		return errors.Wrap(err, "update device last-seen and dr error")
	}

	return nil
}

func updateDeviceActivation(ctx *uplinkContext) error {
	da := ctx.uplinkDataReq.DeviceActivationContext

	// nothing to do when there is no new device activation context
	if da == nil {
		return nil
	}

	// key envelope must not be nil
	if da.AppSKey == nil {
		return errors.New("AppSKey must not be nil")
	}

	appSKey, err := unwrapASKey(da.AppSKey)
	if err != nil {
		return errors.Wrap(err, "unwrap AppSKey error")
	}

	var devAddr lorawan.DevAddr
	copy(devAddr[:], da.DevAddr)

	// if DevAddr and AppSKey are equal, there is nothing to do
	if ctx.device.DevAddr == devAddr && ctx.device.AppSKey == appSKey {
		return nil
	}

	ctx.device.DevAddr = devAddr
	ctx.device.AppSKey = appSKey

	if err := ctx.handler.UpdateDeviceActivation(ctx.ctx, ctx.device.DevEUI, ctx.device.DevAddr, ctx.device.AppSKey); err != nil {
		return errors.Wrap(err, "update device activation error")
	}

	pl := pb.JoinEvent{
		ApplicationId:   uint64(ctx.device.ApplicationID),
		ApplicationName: ctx.application.Name,
		DevEui:          ctx.device.DevEUI[:],
		DeviceName:      ctx.device.Name,
		DevAddr:         ctx.device.DevAddr[:],
		RxInfo:          ctx.uplinkDataReq.RxInfo,
		TxInfo:          ctx.uplinkDataReq.TxInfo,
		Dr:              ctx.uplinkDataReq.Dr,
		Tags:            make(map[string]string),
	}

	// set tags
	for k, v := range ctx.device.Tags.Map {
		if v.Valid {
			pl.Tags[k] = v.String
		}
	}

	vars := make(map[string]string)
	for k, v := range ctx.device.Variables.Map {
		if v.Valid {
			vars[k] = v.String
		}
	}

	err = integration.ForApplicationID(ctx.device.ApplicationID, ctx.gIntegrations).HandleJoinEvent(ctx.ctx, vars, pl)
	if err != nil {
		return errors.Wrap(err, "send join notification error")
	}

	return nil
}

func decryptPayload(ctx *uplinkContext) error {
	var err error

	ctx.data, err = lorawan.EncryptFRMPayload(ctx.device.AppSKey, true, ctx.device.DevAddr, ctx.uplinkDataReq.FCnt, ctx.uplinkDataReq.Data)
	if err != nil {
		return errors.Wrap(err, "decrypt payload error")
	}
	return nil
}

func handleApplicationLayers(ctx *uplinkContext) error {
	// TODO: make application layer configurable
	// * make ports configurable
	// * disable application-layer features
	if ctx.uplinkDataReq.FPort != 200 && ctx.uplinkDataReq.FPort != 201 && ctx.uplinkDataReq.FPort != 202 {
		return nil
	}

	return ctx.handler.Tx(context.Background(), func(context context.Context, handler *store.Handler) error {
		switch ctx.uplinkDataReq.FPort {
		case 200:
			if err := multicastsetup.HandleRemoteMulticastSetupCommand(ctx.ctx, handler, ctx.device.DevEUI, ctx.data); err != nil {
				return errors.Wrap(err, "handle remote multicast setup command error")
			}
		case 201:
			if err := fragmentation.HandleRemoteFragmentationSessionCommand(ctx.ctx, handler, ctx.device.DevEUI, ctx.data); err != nil {
				return errors.Wrap(err, "handle remote fragmentation session command error")
			}
		case 202:
			var timeSinceGPSEpoch time.Duration
			var timeField time.Time
			var err error

			for _, rxInfo := range ctx.uplinkDataReq.RxInfo {
				if rxInfo.TimeSinceGpsEpoch != nil {
					timeSinceGPSEpoch, err = ptypes.Duration(rxInfo.TimeSinceGpsEpoch)
					if err != nil {
						log.WithError(err).Error("time since gps epoch to duration error")
						continue
					}
				} else if rxInfo.Time != nil {
					timeField, err = ptypes.Timestamp(rxInfo.Time)
					if err != nil {
						log.WithError(err).Error("time to timestamp error")
						continue
					}
				}
			}

			// fallback on time field when time since GPS epoch is not available
			if timeSinceGPSEpoch == 0 {
				// fallback on current server time when time field is not available
				if timeField.IsZero() {
					timeField = time.Now()
				}
				timeSinceGPSEpoch = gps.Time(timeField).TimeSinceGPSEpoch()
			}

			if err := clocksync.HandleClockSyncCommand(ctx.ctx, handler, ctx.device.DevEUI, timeSinceGPSEpoch, ctx.data); err != nil {
				return errors.Wrap(err, "handle clocksync command error")
			}
		}
		return nil
	})
}

func handleCodec(ctx *uplinkContext) error {
	codecType := ctx.application.PayloadCodec
	decoderScript := ctx.application.PayloadDecoderScript

	if ctx.deviceProfile.PayloadCodec != "" {
		codecType = ctx.deviceProfile.PayloadCodec
		decoderScript = ctx.deviceProfile.PayloadDecoderScript
	}

	if codecType == codec.None {
		return nil
	}

	start := time.Now()
	b, err := codec.BinaryToJSON(codec.Type(codecType), uint8(ctx.uplinkDataReq.FPort), ctx.device.Variables, decoderScript, ctx.data)
	if err != nil {
		log.WithFields(log.Fields{
			"codec":          codecType,
			"application_id": ctx.device.ApplicationID,
			"f_port":         ctx.uplinkDataReq.FPort,
			"f_cnt":          ctx.uplinkDataReq.FCnt,
			"dev_eui":        ctx.device.DevEUI,
		}).WithError(err).Error("decode payload error")

		errEvent := pb.ErrorEvent{
			ApplicationId:   uint64(ctx.device.ApplicationID),
			ApplicationName: ctx.application.Name,
			DeviceName:      ctx.device.Name,
			DevEui:          ctx.device.DevEUI[:],
			Type:            pb.ErrorType_UPLINK_CODEC,
			Error:           err.Error(),
			FCnt:            ctx.uplinkDataReq.FCnt,
			Tags:            make(map[string]string),
		}

		for k, v := range ctx.device.Tags.Map {
			if v.Valid {
				errEvent.Tags[k] = v.String
			}
		}

		vars := make(map[string]string)
		for k, v := range ctx.device.Variables.Map {
			if v.Valid {
				vars[k] = v.String
			}
		}

		if err := integration.ForApplicationID(ctx.device.ApplicationID, ctx.gIntegrations).HandleErrorEvent(ctx.ctx, vars, errEvent); err != nil {
			log.WithError(err).Error("send error event to integration error")
		}
	}

	log.WithFields(log.Fields{
		"application_id": ctx.application.ID,
		"codec":          codecType,
		"duration":       time.Since(start),
	}).Debug("payload codec completed Decode execution")

	ctx.objectJSON = string(b)

	return nil
}

func handleIntegrations(ctx *uplinkContext) error {
	pl := pb.UplinkEvent{
		ApplicationId:   uint64(ctx.device.ApplicationID),
		ApplicationName: ctx.application.Name,
		DeviceName:      ctx.device.Name,
		DevEui:          ctx.device.DevEUI[:],
		RxInfo:          ctx.uplinkDataReq.RxInfo,
		TxInfo:          ctx.uplinkDataReq.TxInfo,
		Dr:              ctx.uplinkDataReq.Dr,
		Adr:             ctx.uplinkDataReq.Adr,
		FCnt:            ctx.uplinkDataReq.FCnt,
		FPort:           ctx.uplinkDataReq.FPort,
		Data:            ctx.data,
		ObjectJson:      ctx.objectJSON,
		Tags:            make(map[string]string),
		ConfirmedUplink: ctx.uplinkDataReq.ConfirmedUplink,
		DevAddr:         ctx.device.DevAddr[:],
	}

	// set tags
	for k, v := range ctx.device.Tags.Map {
		if v.Valid {
			pl.Tags[k] = v.String
		}
	}
	vars := make(map[string]string)
	for k, v := range ctx.device.Variables.Map {
		if v.Valid {
			vars[k] = v.String
		}
	}

	bgCtx := context.Background()
	bgCtx = context.WithValue(bgCtx, logging.ContextIDKey, ctx.ctx.Value(logging.ContextIDKey))

	// Handle the actual integration handling in a Go-routine so that the
	// as.HandleUplinkData api can return.
	go func() {
		err := integration.ForApplicationID(ctx.device.ApplicationID, ctx.gIntegrations).HandleUplinkEvent(bgCtx, vars, pl)
		if err != nil {
			log.WithError(err).Error("send uplink event error")
		}
	}()

	return nil
}

func unwrapASKey(ke *common.KeyEnvelope) (lorawan.AES128Key, error) {
	var key lorawan.AES128Key

	if ke.KekLabel == "" {
		copy(key[:], ke.AesKey)
		return key, nil
	}

	joinServer := joinserver.GetSettings()
	for i := range joinServer.KEK.Set {
		if joinServer.KEK.Set[i].Label == ke.KekLabel {
			kek, err := hex.DecodeString(joinServer.KEK.Set[i].KEK)
			if err != nil {
				return key, errors.Wrap(err, "decode kek error")
			}

			block, err := aes.NewCipher(kek)
			if err != nil {
				return key, errors.Wrap(err, "new cipher error")
			}

			b, err := keywrap.Unwrap(block, ke.AesKey)
			if err != nil {
				return key, errors.Wrap(err, "key unwrap error")
			}

			copy(key[:], b)
			return key, nil
		}
	}

	return key, fmt.Errorf("unknown kek label: %s", ke.KekLabel)
}
