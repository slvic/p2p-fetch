package configs

import (
	"flag"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type AppConfig struct {
	Excel Excel `hcl:"excel"`
}

type Excel struct {
	FileName string `hcl:"file_name"`
}

func GetConfig(fileName string) (AppConfig, error) {
	var appConfig AppConfig
	err := hclsimple.DecodeFile(fileName, nil, appConfig)
	if err != nil {
		return AppConfig{}, fmt.Errorf("failed to load configuration: %s", err.Error())
	}

	flag.StringVar(
		&appConfig.Excel.FileName,
		"C",
		appConfig.Excel.FileName,
		"excel working file name")

	if len(appConfig.Excel.FileName) == 0 {
		return AppConfig{}, fmt.Errorf("please specify excel working file name")
	}

	return appConfig, nil
}
