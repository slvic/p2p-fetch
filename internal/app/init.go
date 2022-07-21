package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/slvic/p2p-fetch/internal/configs"
)

const (
	defaulConfigPath = "../../configs/config.hcl"
)

type App struct {
}

func Initialize(ctx context.Context) (*App, error) {
	config, err := configs.GetConfig(defaulConfigPath)
	if err != nil {
		return nil, fmt.Errorf("could not get config: %s", err.Error())
	}

	flag.StringVar(
		&config.Binance.Address,
		"a",
		config.Binance.Address,
		"binance api full address")

	flag.Parse()
	return &App{}, nil
}

func (a *App) Run(ctx context.Context) error {
	return nil
}
