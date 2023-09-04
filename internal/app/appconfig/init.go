package appconfig

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var Conf *appConfig

func init() {
	viperConfig := viper.New()
	loadBase(viperConfig)
	mergeConfig(viperConfig)
	c := Config{}
	err := viperConfig.Unmarshal(&c, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	})
	if err != nil {
		panic(err)
	}
	Conf = &appConfig{
		Viper:  viperConfig,
		Config: c,
	}
}

func loadBase(viperConfig *viper.Viper) {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	baseConfig := filepath.Join(basePath, "configs/app.yaml")
	viperConfig.SetConfigFile(baseConfig)
	err = viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func mergeConfig(viperConfig *viper.Viper) {
	var configFile string
	envConfigPath := os.Getenv("APP_CONFIG_PATH")
	if envConfigPath != "" {
		configFile = envConfigPath
	}
	if configFile == "" && len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	viperConfig.SetConfigFile(configFile)
	err := viperConfig.MergeInConfig()
	if err != nil {
		panic(err)
	}
}
