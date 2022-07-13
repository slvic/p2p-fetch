package config

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	_ "github.com/hashicorp/hcl/v2/hclsimple"
	"log"
)

type Binance struct {
	Address string `hcl:"address"`
}

func GetConfig(filePath string) Binance {
	var config Binance
	err := hclsimple.DecodeFile(filePath, nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Print("binance configuration parsed successfully")
	return config
}
