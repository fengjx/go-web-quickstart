package env

import (
	"log"
	"os"
	"path/filepath"
)

var AppName string
var AppPath string

func init() {
	var err error
	AppPath, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	AppName = filepath.Base(os.Args[0])
	log.Println("app name", AppName)
	log.Println("app path", AppPath)
}

type ENV string

const (
	Local ENV = "local"
	Dev   ENV = "dev"
	Test  ENV = "test"
	Prod  ENV = "prod"
)

func GetEnv() ENV {
	env := os.Getenv("APP_ENV")
	switch ENV(env) {
	case Test:
		return Test
	case Prod:
		return Prod
	case Dev:
		return Dev
	default:
		return Local
	}
}

func IsProd() bool {
	return GetEnv() == Prod
}

func IsTest() bool {
	return GetEnv() == Test
}

func IsDev() bool {
	return GetEnv() == Dev
}

func IsLocal() bool {
	return GetEnv() == Local
}
