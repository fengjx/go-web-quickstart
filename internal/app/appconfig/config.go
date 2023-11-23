package appconfig

import "github.com/spf13/viper"

type AppConfig struct {
	*viper.Viper
	Config
}

type Config struct {
	Env    string                  `yaml:"envh"`
	Server ServerConfig            `yaml:"server"`
	DB     map[string]*DbConfig    `yaml:"db"`
	Redis  map[string]*RedisConfig `yaml:"redis"`
}

type ServerConfig struct {
	Listen     string   `yaml:"listen"`
	Template   []string `yaml:"template"`
	Static     string   `yaml:"static"`
	AuthSecret string   `yaml:"auth-secret"`
	Accesslog  string   `yaml:"accesslog"`
}

type DbConfig struct {
	Type    string `yaml:"type"`
	Dsn     string `yaml:"dsn"`
	MaxIdle int    `yaml:"max-idle"`
	MaxConn int    `yaml:"max-conn"`
	ShowSQL bool   `yaml:"show-sql"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
