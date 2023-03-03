package applog

import "fengjx/go-web-quickstart/pkg/logger"

var Log logger.Logger

func Init() {
	Log = logger.New()
}
