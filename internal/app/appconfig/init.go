package appconfig

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/env"
)

var conf *AppConfig

func Init() {
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
	conf = &AppConfig{
		Viper:  viperConfig,
		Config: c,
	}
}

func loadBase(viperConfig *viper.Viper) {
	configName := "configs/app.yaml"
	applog.Log.Infof("load config: %s", configName)
	viperConfig.SetConfigFile(configName)
	err := viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func mergeConfig(viperConfig *viper.Viper) {
	var configFile string
	envConfigPath := os.Getenv("APP_CONFIG")
	if envConfigPath != "" {
		configFile = envConfigPath
	}
	if configFile == "" && len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	if configFile == "" {
		appEnv := env.GetEnv()
		configFile = path.Join("configs", fmt.Sprintf("app-%s.yaml", appEnv))
	}
	applog.Log.Infof("merge config: %s", configFile)
	viperConfig.SetConfigFile(configFile)
	err := viperConfig.MergeInConfig()
	if err != nil {
		panic(err)
	}
}

func GetConfig() *AppConfig {
	return conf
}
