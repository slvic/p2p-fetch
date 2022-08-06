package api

import (
	"github.com/slvic/p2p-fetch/internal/configs"
)

type Bestchange struct {
	config configs.Bestchange
}

func NewBestchangeParser(cfg configs.Bestchange) *Bestchange {
	return &Bestchange{cfg}
}

func (b Bestchange) GetData() {
	err := getBcApiFile(b.config.ApiUrl)
	if err != nil {
		return
	}

	err = unzipSource(bcApiZipFileName, bcApiFolder)
	if err != nil {
		return
	}

	currencies, err := getRawCurrencies(currenciesFile)
	if err != nil {
		return
	}
	exchangers, err := getRawExchangers(exchangerOfficesFile)
	if err != nil {
		return
	}
	rawExchangeRates, err := getRawExchangeRates(exchangeRatesFile)
	if err != nil {
		return
	}

	exchangeRates := getExchangeRates(rawExchangeRates, exchangers, currencies)

	for _, exchangeRate := range exchangeRates {

	}

}
