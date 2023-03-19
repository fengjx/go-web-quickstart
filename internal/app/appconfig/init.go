package appconfig

import (
	"fmt"
	"os"
)

var Conf *Config

func init() {
	var configFile string
	envConfigPath := os.Getenv("APP_CONFIG_PATH")
	if envConfigPath != "" {
		configFile = envConfigPath
	}
	if configFile == "" && len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	c, err := initConfig(configFile)
	if err != nil {
		info := fmt.Sprintf("Load config err - %s, custom config file: %s", err.Error(), configFile)
		panic(info)
	}
	Conf = c
}

func Get() *Config {
	return Conf
}
