package configs

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type AppConfig struct {
	App        App        `hcl:"app,block"`
	Binance    Binance    `hcl:"binance,block"`
	Bestchange Bestchange `hcl:"bestchange,block"`
}

type App struct {
	FetchInterval int64 `hcl:"fetchInterval"`
}

type Binance struct {
	Address string   `hcl:"address"`
	Assets  []string `hcl:"assets"`
	Fiats   []string `hcl:"fiats"`
}

type Bestchange struct {
	BaseUrl string `hcl:"baseurl"`
}

func GetConfig(fileName string) (AppConfig, error) {
	var appConfig AppConfig
	err := hclsimple.DecodeFile(fileName, nil, &appConfig)
	if err != nil {
		return AppConfig{}, fmt.Errorf("failed to load configuration: %s", err.Error())
	}

	return appConfig, nil
}
