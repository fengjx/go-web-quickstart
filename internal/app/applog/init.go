package applog

import (
	"log"
	"math"

	"github.com/fengjx/go-halo/logger"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/hook"
)

var Log logger.Logger

func Init() {
	logConfig := appconfig.Conf.Log
	if logConfig.Appender == "console" {
		Log = logger.NewConsole()
	} else {
		if logConfig.OpenTrace {
			Log = logger.New(logConfig.Level, logConfig.Path, logConfig.MaxSize, logConfig.MaxDays, logger.WithTrace())
		} else {
			Log = logger.New(logConfig.Level, logConfig.Path, logConfig.MaxSize, logConfig.MaxDays)
		}
	}
	Log.Infof("app log init")
	hook.AddStopHook(Flush, math.MaxInt)
}

func Flush() {
	logConfig := appconfig.Conf.Config.Log
	if Log != nil && logConfig.Appender != "console" {
		log.Println("flush log")
		Log.Warn("flush log")
		Log.Flush()
	}
}
