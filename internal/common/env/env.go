package env

import "fengjx/go-web-quickstart/internal/app/appconfig"

func IsProd() bool {
	return appconfig.GetEnv() == "prod"
}

func IsTest() bool {
	return appconfig.GetEnv() == "test"
}

func IsDev() bool {
	return appconfig.GetEnv() == "dev"
}
