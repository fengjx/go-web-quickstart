package applog

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fengjx/go-halo/hook"
	"github.com/fengjx/go-halo/logger"

	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/common/env"
)

var Log logger.Logger
var LogPath string

func init() {
	if env.IsDev() || env.IsLocal() {
		Log = logger.NewConsole()
		return
	}
	appPath := env.AppPath
	app := filepath.Base(os.Args[0])
	LogPath = filepath.Join(appPath, "logs")
	err := os.MkdirAll(LogPath, 0644)
	if err != nil {
		panic(err)
	}
	logfile := filepath.Join(LogPath, app+".log")
	Log = logger.New(logger.InfoLevel, logfile, 100, 3)
	log.Println("log file", logfile)
	Log.Infof("logfile: %s", logfile)
	hook.AddCustomStartHook(common.HookEventShutdown, func() {
		Log.Flush()
	}, common.HookOrderLowest)
}
