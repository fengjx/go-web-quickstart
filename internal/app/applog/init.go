package applog

import (
	"context"
	"log"
	"math"
	"net/http"

	"github.com/fengjx/go-halo/logger"
	"github.com/google/uuid"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/hook"
)

var Log logger.Logger

func Init() {
	logConfig := appconfig.Conf.Log
	if logConfig.Appender == "console" {
		Log = logger.NewConsole()
	} else {
		Log = logger.New(logConfig.Level, logConfig.Path, logConfig.MaxSize, logConfig.MaxDays)
	}
	Log.Infof("app log init")
	hook.AddStopHook(Flush, math.MaxInt)
}

func WithRequest(ctx context.Context, req *http.Request) context.Context {
	id := getRequestID(req)
	if id == "" {
		id = uuid.New().String()
	}
	ctx = context.WithValue(ctx, logger.RequestIDKey, id)
	return ctx
}

func getRequestID(req *http.Request) string {
	return req.Header.Get(logger.RequestIDKey)
}

func Flush() {
	log.Println("flush log")
	logConfig := appconfig.Conf.Config.Log
	if Log != nil && logConfig.Appender != "console" {
		Log.Sync()
	}
}
