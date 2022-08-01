package app

import (
	"context"
	"fmt"
	"github.com/slvic/p2p-fetch/internal/configs"
	"github.com/slvic/p2p-fetch/pkg/bestchange"
	"github.com/slvic/p2p-fetch/pkg/markets/binance"
	"log"
)

const (
	defaulConfigPath = "../../configs/config.hcl"
)

type App struct {
	bestChange *bestchange.Bestchange
	binance    *binance.Binance
	config     configs.App
}

func Initialize(ctx context.Context) (*App, error) {
	config, err := configs.GetConfig(defaulConfigPath)
	if err != nil {
		return nil, fmt.Errorf("could not get config: %s", err.Error())
	}

	bestchangeParser := bestchange.New(config.Bestchange)
	binanceApi := binance.New(config.Binance)

	return &App{
		bestChange: bestchangeParser,
		binance:    binanceApi,
		config:     config.App,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		assets, err := a.bestChange.GetAssets()
		if err != nil {
			log.Printf("could not get assets: %s", err.Error())
			return
		}
		err = a.bestChange.GetExchangers(assets)
		if err != nil {
			log.Printf("could not get exchangers: %s", err.Error())
			return
		}
	}()

	go func() {
		a.binance.GetData()
	}()
}
