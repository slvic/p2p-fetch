package configs

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type AppConfig struct {
	Binance    Binance    `hcl:"binance"`
	Bestchange Bestchange `hcl:"bestchange"`
}

type Binance struct {
	Address string `hcl:"address"`
}

type Bestchange struct {
	BaseUrl string `hcl:"baseurl"`
}

func GetConfig(fileName string) (AppConfig, error) {
	var appConfig AppConfig
	err := hclsimple.DecodeFile(fileName, nil, appConfig)
	if err != nil {
		return AppConfig{}, fmt.Errorf("failed to load configuration: %s", err.Error())
	}

	return appConfig, nil
}
