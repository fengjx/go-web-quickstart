package appconfig

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name     string
	Env      string
	BasePath string
	Server   *ServerConfig
	DB       map[string]*DbConfig
	Redis    map[string]*RedisConfig
	Kv       map[string]string
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

func initConfig(configFile string) (*Config, error) {
	c := new(Config)
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	baseConfig := filepath.Join(basePath, "configs/app.yaml")
	err = load(c, baseConfig)
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
	c.BasePath = basePath
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
