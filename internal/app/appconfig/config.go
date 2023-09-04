package appconfig

import "github.com/spf13/viper"

type appConfig struct {
	*viper.Viper
	Config
}

type Config struct {
	Name     string                  `yaml:"name"`
	Env      string                  `yaml:"env"`
	BasePath string                  `yaml:"base-path"`
	Server   ServerConfig            `yaml:"server"`
	DB       map[string]*DbConfig    `yaml:"db"`
	Redis    map[string]*RedisConfig `yaml:"redis"`
	Log      LogConfig               `yaml:"log"`
}

type LogConfig struct {
	Appender  string `yaml:"appender"`
	Level     string `yaml:"level"`
	Path      string `yaml:"path"`
	MaxSize   int    `yaml:"max-size"`
	MaxDays   int    `yaml:"max-days"`
	OpenTrace bool   `yaml:"open-trace"`
}

type ServerConfig struct {
	Host       string   `yaml:"host"`
	Port       int      `yaml:"port"`
	Access     string   `yaml:"access"`
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

func GetEnv() string {
	return Conf.Config.Env
}
