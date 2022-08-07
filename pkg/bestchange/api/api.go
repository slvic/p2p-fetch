package api

import (
	"context"
	"fmt"
	"github.com/mehanizm/iuliia-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/slvic/p2p-fetch/internal/configs"
	"github.com/slvic/p2p-fetch/pkg/bestchange/models"
	"golang.org/x/sync/errgroup"
	"log"
	"strings"
	"time"
)

type Bestchange struct {
	config configs.Bestchange
}

func NewBestchangeParser(cfg configs.Bestchange) *Bestchange {
	return &Bestchange{cfg}
}

func (b Bestchange) GetData(ctx context.Context) {
	err := getBcApiFile(b.config.ApiUrl)
	if err != nil {
		log.Printf("could not get bestchange api file: %s", err.Error())
		return
	}

	err = unzipSource(bcApiZipFileName, bcApiFolder)
	if err != nil {
		log.Printf("could not unzip bestchange api file: %s", err.Error())
		return
	}

	rawGetter, _ := errgroup.WithContext(ctx)

	rawCurrencies := make(chan map[int]string, 1)
	rawExchangers := make(chan map[int]string, 1)
	rawExchangeRates := make(chan []models.RawExchangeRate, 1)

	rawGetter.Go(func() error {
		currencies, err := getRawCurrencies(currenciesFile)
		if err != nil {
			return fmt.Errorf("could not get raw currencies: %w", err)
		}
		rawCurrencies <- currencies
		return nil
	})
	rawGetter.Go(func() error {
		exchangers, err := getRawExchangers(exchangerOfficesFile)
		if err != nil {
			return fmt.Errorf("could not get raw exchangers: %w", err)
		}
		rawExchangers <- exchangers
		return nil
	})
	rawGetter.Go(func() error {
		exchangeRates, err := getRawExchangeRates(exchangeRatesFile)
		if err != nil {
			return fmt.Errorf("could not get raw exchange rates: %w", err)
		}
		rawExchangeRates <- exchangeRates
		return nil
	})

	if err = rawGetter.Wait(); err != nil {
		log.Printf("could not get raw bestchange data: %s", err.Error())
		return
	}

	exchangeRates := getExchangeRates(<-rawExchangeRates, <-rawExchangers, <-rawCurrencies)

	replacer := strings.NewReplacer(" ", "_", "-", "_", "(", "", ")", "", "/", "", ".", "")
	index := 0
	for _, exchangeRate := range exchangeRates {
		index += 1
		{ //get rate
			prometheus.NewSummary(prometheus.SummaryOpts{
				Namespace: "bestchange",
				Name: fmt.Sprintf(
					"%s_exchanger_%s_source_to_%s_target_index_%d",
					replacer.Replace(iuliia.Wikipedia.Translate(exchangeRate.ExchangerName)),
					replacer.Replace(iuliia.Wikipedia.Translate(exchangeRate.SourceCurrency)),
					replacer.Replace(iuliia.Wikipedia.Translate(exchangeRate.TargetCurrency)),
					index,
				),
			}).Observe(exchangeRate.GetRate)
		}
	}
	log.Printf("bestchange api data is successfully gathered: %v", time.Now())
}
