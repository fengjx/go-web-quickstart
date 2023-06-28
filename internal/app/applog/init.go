package applog

import "github.com/fengjx/go-web-quickstart/pkg/logger"

var Log logger.Logger

func Init() {
	Log = logger.New()
	Log.Infof("app log init")
}
