package appconfig

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Conf *appConfig

func init() {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	baseConfig := filepath.Join(basePath, "configs/app.yaml")

	config := viper.New()
	config.SetConfigFile(baseConfig)
	err = config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var configFile string
	envConfigPath := os.Getenv("APP_CONFIG_PATH")
	if envConfigPath != "" {
		configFile = envConfigPath
	}
	if configFile == "" && len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	config.SetConfigFile(configFile)
	err = config.MergeInConfig()
	if err != nil {
		panic(err)
	}

	c := Config{}
	err = config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	Conf = &appConfig{
		Viper:  config,
		Config: c,
	}
}
