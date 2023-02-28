package app

import (
	"fengjx/go-web-quickstart/internal/app/config"
	"fengjx/go-web-quickstart/pkg/logger"
	"fmt"
	"os"
)

var Log logger.Logger

var Config *config.Config

func init() {
	Log = logger.New()
	var configFile string
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	c, err := config.New(configFile)
	if err != nil {
		info := fmt.Sprintf("Load config err - %s, custom config file: %s", err.Error(), configFile)
		panic(info)
	}
	Config = c
}

func NewServer() Server {
	var serv Server = &ginServer{}
	return serv
}

func IsProd() bool {
	return Config.Env == "prod"
}

func IsTest() bool {
	return Config.Env == "test"
}

func IsDev() bool {
	return Config.Env == "dev"
}

func GetProp(key string) string {
	return Config.Kv[key]
}
