package appconfig

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Name   string
	Env    string
	Server *ServerConfig
	DB     map[string]*DbConfig
	Redis  map[string]*RedisConfig
	Kv     map[string]string
}

type ServerConfig struct {
	Host     string
	Port     int
	Access   string
	Template []string
}

type DbConfig struct {
	Type    string
	Dsn     string
	MaxIdle int  `yaml:"max-idle"`
	MaxConn int  `yaml:"max-conn"`
	ShowSQL bool `yaml:"show-sql"`
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func initConfig(configFile string) (*Config, error) {
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

func GetEnv() string {
	return Conf.Env
}

func GetProp(key string) string {
	return Conf.Kv[key]
}
