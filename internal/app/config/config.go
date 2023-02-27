package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name   string
	Env    string
	Server serverConfig
	Mysql  mysqlConfig
	Redis  redisConfig
}

type serverConfig struct {
	Host     string
	Port     int
	Access   string
	Template []string
}

type mysqlConfig struct {
}

type redisConfig struct {
}

func New(configFile string) (*Config, error) {
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	c := new(Config)
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}
	return c, err
}
