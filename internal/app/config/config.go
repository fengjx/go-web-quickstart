package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Name   string
	Env    string
	Server serverConfig
	Mysql  mysqlConfig
	Redis  redisConfig
	Kv     map[string]string
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
	c := new(Config)
	err := load(c, "configs/app.yaml")
	if err != nil {
		return nil, err
	}
	if configFile == "" {
		return c, nil
	}
	// merge
	err = load(c, configFile)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func load(c *Config, configFile string) error {
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return err
	}
	return err
}
