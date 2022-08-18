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
	"runtime"
	"sync"
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
	ctx, cancelFunc := context.WithCancel(ctx)
	go startMetricsGatherer(cancelFunc)

	log.Printf("\napp is running...\n")
	dur := time.Duration(a.config.FetchInterval) * time.Hour
	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	oomTicker := time.NewTicker(time.Second)
	defer oomTicker.Stop()

	go func(ctx context.Context) {
		var m runtime.MemStats
		select {
		case <-oomTicker.C:
			runtime.ReadMemStats(&m)
			// For info on each, see: https://golang.org/pkg/runtime/#MemStats
			fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
			fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
			fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
			fmt.Printf("\tNumGC = %v\n", m.NumGC)
		case <-ctx.Done():
			return
		}
	}(ctx)

	a.gatherData(ctx)
outerLoop:
	for {
		select {
		case <-ticker.C:
			a.gatherData(ctx)
		case <-ctx.Done():
			break outerLoop
		}
	}

	return nil
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func startMetricsGatherer(cancel context.CancelFunc) {
	r := http.NewServeMux()
	r.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("could not start a metrics gatherer: %s", err.Error())
		cancel()
	}
}

func (a *App) gatherData(ctx context.Context) {
	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		a.bestchange.GetData(ctx)
		wg.Done()
	}()

	go func() {
		a.binance.GetAllData(ctx)
		wg.Done()
	}()

	wg.Wait()
	log.Printf("all data is successfully fetched, next fetch will start in %s", startTime.Add(4*time.Hour))
}
