package binance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/slvic/p2p-fetch/internal/configs"
	"github.com/slvic/p2p-fetch/pkg/markets/models"
	"io/ioutil"
	"log"
	"net/http"
)

type Binance struct {
	config configs.Binance
}

func New(cfg configs.Binance) *Binance {
	return &Binance{config: cfg}
}

func getOptions(asset, fiat string) models.BinanceRequest {
	return models.BinanceRequest{
		Asset:         asset,
		Fiat:          fiat,
		MerchantCheck: true,
		Page:          1,
		PayTypes:      []string{"BANK"},
		PublisherType: nil,
		Rows:          20,
		TradeType:     "SELL",
	}
}

func (b *Binance) GetAllData() {
	for _, asset := range b.config.Assets {
		for _, fiat := range b.config.Fiats {
			options := getOptions(asset, fiat)
			err := b.getData(&options)
			if err != nil {
				log.Printf("could not get binance data: %s", err.Error())
			}
		}
	}
}

func (b *Binance) getData(options *models.BinanceRequest) error {
	var binanceResponse models.BinanceResponse

	response, err := b.sendRequest(options)
	if err != nil {
		return fmt.Errorf("could not send request: %s", err.Error())
	}

	err = json.Unmarshal(response, &binanceResponse)
	if err != nil {
		return fmt.Errorf("could not unmarshal responce body: %s", err.Error())
	}

	return nil
}

func (b Binance) sendRequest(options *models.BinanceRequest) ([]byte, error) {
	bodyBytes, err := json.Marshal(&options)
	if err != nil {
		return nil, fmt.Errorf("could not marshal options: %s", err.Error())
	}
	bodyReader := bytes.NewReader(bodyBytes)

	response, err := http.Post(b.config.Address, "application/json", bodyReader)
	if err != nil {
		return nil, fmt.Errorf("could not send a request: %s", err.Error())
	}

	responseBodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read a responce body: %s", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unsuccessfull request, status code %d, response body: %s",
			response.StatusCode,
			string(responseBodyBytes))
	}

	return responseBodyBytes, nil
}
