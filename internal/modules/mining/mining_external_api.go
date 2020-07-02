package mining

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/brocaar/lorawan"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"

	api "github.com/mxc-foundation/lpwan-app-server/api/m2m-serves-appserver"
	"github.com/mxc-foundation/lpwan-app-server/internal/backend/m2m_client"
	"github.com/mxc-foundation/lpwan-app-server/internal/config"
	"github.com/mxc-foundation/lpwan-app-server/internal/modules/gateway"
	"github.com/mxc-foundation/lpwan-app-server/internal/storage"
)

type MiningAPI struct {
}

var (
	miningAPI MiningAPI

	dailyMxcPrice, totalMiningValue float64
)

func GetMiningAPI() *MiningAPI {
	return &miningAPI
}

func Setup(conf config.Config) error {
	log.Info("mining cron task begin...")

	// we need to seed it as we use it later
	rand.Seed(time.Now().UnixNano())

	mconf := conf.ApplicationServer.MiningSetUp
	if mconf.MinValue <= 0 || mconf.MinValue > mconf.MaxValue {
		err := fmt.Errorf("invalid mining configuration, min_value %d and max_value %d", mconf.MinValue, mconf.MaxValue)
		log.Error(err)
		return err
	}
	oneMxc, err := miningAPI.getUSDprice(conf)
	if err != nil {
		log.WithError(err).Error("tokenMining/Unable to get USD price from CMC")
		return err
	}
	dailyMxcPrice = oneMxc

	totalMiningValue = miningAPI.calcTotalMiningValue(conf, dailyMxcPrice)

	c := cron.New()
	exeTime := config.C.ApplicationServer.MiningSetUp.ExecuteTime

	err = c.AddFunc(exeTime, func() {
		log.Info("Start token mining")
		go func() {
			err := miningAPI.tokenMining(context.Background(), conf)
			if err != nil {
				log.WithError(err).Error("tokenMining Error")
			}
		}()
	})
	if err != nil {
		log.WithError(err).Error("Start mining cron task failed")
	}
	go c.Start()

	priceCron := cron.New()
	// update MXC real price everyday 3 am
	err = priceCron.AddFunc("0 0 3 * * ?", func() {
		log.Info("Get new MXC Price")
		go func() {
			oneMxc, err := miningAPI.getUSDprice(conf)
			if err != nil {
				log.WithError(err).Error("tokenMining/Unable to get USD price from CMC")
			}
			dailyMxcPrice = oneMxc

			totalMiningValue = miningAPI.calcTotalMiningValue(conf, dailyMxcPrice)
		}()
	})
	if err != nil {
		log.WithError(err).Error("Start mining cron task failed")
	}
	go priceCron.Start()

	return nil
}

func (a *MiningAPI) calcTotalMiningValue(conf config.Config, mxcPrice float64) float64 {
	mconf := conf.ApplicationServer.MiningSetUp
	randUSD := float64(rand.Int63n(mconf.MaxValue-mconf.MinValue) + mconf.MinValue)
	return mxcPrice * randUSD
}

// getUSDprice returns 1 USD price in MXC
func (a *MiningAPI) getUSDprice(conf config.Config) (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/tools/price-conversion", nil)
	if err != nil {
		log.WithError(err).Error("CMC client error")
		os.Exit(1)
	}

	q := url.Values{}
	//q.Add("id", "2")
	q.Add("symbol", "USD")
	q.Add("amount", "1")
	q.Add("convert", "MXC")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", conf.ApplicationServer.MiningSetUp.CMCKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("CMC request error")
		os.Exit(1)
	}

	if resp.Status == "200 OK" {

		respBody, _ := ioutil.ReadAll(resp.Body)

		var cmc CMC
		err = json.Unmarshal(respBody, &cmc)
		if err != nil {
			log.Println("JSON unmarshal error: ", err)
		}

		return cmc.Data.Quote.MXC.Price, nil
	}

	err = errors.New("getUSDprice/Get USD price from cmc error")
	return 0, err
}

// getMXCprice returns amount of MXC in USD
func (a *MiningAPI) GetMXCprice(conf config.Config, amount string) (price float64, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/tools/price-conversion", nil)
	if err != nil {
		log.WithError(err).Error("CMC client error")
		os.Exit(1)
	}

	q := url.Values{}
	//q.Add("id", "2")
	q.Add("symbol", "MXC")
	q.Add("amount", amount)
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", conf.ApplicationServer.MiningSetUp.CMCKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("CMC request error")
		os.Exit(1)
	}

	if resp.Status == "200 OK" {
		respBody, _ := ioutil.ReadAll(resp.Body)

		var cmc CMC
		err = json.Unmarshal(respBody, &cmc)
		if err != nil {
			log.Println("JSON unmarshal error: ", err)
		}

		return cmc.Data.Quote.USD.Price, nil
	}

	err = errors.New("GetMXCprice/Unable to get the MXC price from cmc")
	return 0, err
}

func (a *MiningAPI) tokenMining(ctx context.Context, conf config.Config) error {
	currentTime := time.Now().Unix()

	// get the gateway list that should receive the mining tokens
	miningGws, err := gateway.GetGatewayAPI().GetGatewayMiningList(ctx, currentTime, conf.ApplicationServer.MiningSetUp.GwOnlineLimit)
	if err != nil {
		if err == storage.ErrDoesNotExist {
			log.Info("No gateway online longer than 24 hours")
			return nil
		}

		log.WithError(err).Error("Cannot get mining gateway list from DB.")
		return err
	}

	if len(miningGws) == 0 {
		return nil
	}

	var macs []string

	// update the first heartbeat = 0
	for _, v := range miningGws {
		//err := storage.UpdateFirstHeartbeat(ctx, storage.DB(), v, currentTime)
		err := gateway.GetGatewayAPI().UpdateFirstHeartbeatToZero(ctx, v)
		if err != nil {
			log.WithError(err).Error("tokenMining/update first heartbeat to zero error")
		}
		mac := lorawan.EUI64.String(v)
		macs = append(macs, mac)
	}

	// 24 hours = 1440 mins
	amount := totalMiningValue / 1440 * 10

	miningSent := false
	// if error, resend after one minute
	for !miningSent {
		err := a.sendMining(ctx, macs, dailyMxcPrice, amount)
		if err != nil {
			log.WithError(err).Error("send mining request to m2m error")
			time.Sleep(60 * time.Second)
			continue
		}
		miningSent = true
	}

	return nil
}

func (a *MiningAPI) sendMining(ctx context.Context, macs []string, mxc_price, amount float64) error {
	m2mClient, err := m2m_client.GetPool().Get(config.C.M2MServer.M2MServer, []byte(config.C.M2MServer.CACert),
		[]byte(config.C.M2MServer.TLSCert), []byte(config.C.M2MServer.TLSKey))
	if err != nil {
		log.WithError(err).Error("create m2mClient for mining error")

		return err
	}

	miningClient := api.NewMiningServiceClient(m2mClient)

	resp, err := miningClient.Mining(ctx, &api.MiningRequest{
		GatewayMac:    macs,
		MiningRevenue: amount,
		MxcPrice:      mxc_price,
	})
	if err != nil {
		log.WithError(err).Error("Mining API request error")
		return err
	}

	// if response == false, resend the request to m2m
	for !resp.Status {
		time.Sleep(60 * time.Second)
		log.Println("Resend mining request.......")
		resp, err = miningClient.Mining(ctx, &api.MiningRequest{
			GatewayMac:    macs,
			MiningRevenue: amount,
			MxcPrice:      mxc_price,
		})
		if err != nil {
			log.WithError(err).Error("Mining API request error")
			return err
		}
	}

	return nil
}
