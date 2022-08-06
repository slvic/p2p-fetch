package app

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slvic/p2p-fetch/internal/configs"
	"github.com/slvic/p2p-fetch/pkg/bestchange/api"
	"github.com/slvic/p2p-fetch/pkg/markets/binance"
	"log"
	"net/http"
	"time"
)

const (
	defaultConfigPath = "configs/config.hcl"
)

type App struct {
	bestchange *api.Bestchange
	binance    *binance.Binance
	config     configs.App
}

func Initialize(ctx context.Context) (*App, error) {
	config, err := configs.GetConfig(defaultConfigPath)
	if err != nil {
		return nil, fmt.Errorf("could not get config: %s", err.Error())
	}

	bestchangeApi := api.NewBestchangeParser(config.Bestchange)
	binanceApi := binance.New(config.Binance)

	return &App{
		bestchange: bestchangeApi,
		binance:    binanceApi,
		config:     config.App,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	err := startMetricsGatherer()
	if err != nil {
		return fmt.Errorf("could not start metrics gatherer: %w", err)
	}

	log.Printf("\napp is running...\n")
	dur := time.Duration(a.config.FetchInterval) * time.Hour
	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	a.gatherData()
outerLoop:
	for {
		select {
		case <-ticker.C:
			a.gatherData()
		case <-ctx.Done():
			break outerLoop
		}
	}

	return nil
}

func startMetricsGatherer() error {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) gatherData() {
	go func() {
		a.bestchange.GetData()
	}()

	go func() {
		a.binance.GetAllData()
	}()
}
