package config

import (
	"github.com/PGo-Projects/output"
	"github.com/spf13/viper"
)

var (
	Filename string
	DevRun   bool
	ProdRun  bool
)

func Init() {
	if Filename != "" {
		viper.SetConfigFile(Filename)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		output.ErrorAndExit(err, 1)
	}

	ProdRun = !DevRun
}
