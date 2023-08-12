package appconfig

import "github.com/spf13/viper"

type appConfig struct {
	*viper.Viper
	Config
}

type Config struct {
	Name     string
	Env      string
	BasePath string
	Server   *ServerConfig
	DB       map[string]*DbConfig
	Redis    map[string]*RedisConfig
	Log      *LogConfig
}

type LogConfig struct {
	Appender string
	Level    string
	Path     string
	MaxSize  int `yaml:"max-size"`
	MaxDays  int `yaml:"max-days"`
}

type ServerConfig struct {
	Host       string
	Port       int
	Access     string
	Template   []string
	AuthSecret string `yaml:"auth-secret"`
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

func GetEnv() string {
	return Conf.Config.Env
}
