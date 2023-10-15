package appconfig

import "github.com/spf13/viper"

type appConfig struct {
	*viper.Viper
	Config
}

type Config struct {
	Server ServerConfig            `yaml:"server"`
	DB     map[string]*DbConfig    `yaml:"db"`
	Redis  map[string]*RedisConfig `yaml:"redis"`
}

type ServerConfig struct {
	Host       string   `yaml:"host"`
	Port       int      `yaml:"port"`
	Template   []string `yaml:"template"`
	AuthSecret string   `yaml:"auth-secret"`
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
