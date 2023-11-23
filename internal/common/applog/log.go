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
var LogPath = filepath.Join("/var", "log")

func init() {
	if env.IsDev() || env.IsLocal() {
		Log = logger.NewConsole()
		return
	}
	err := os.MkdirAll(LogPath, 0644)
	if err != nil {
		panic(err)
	}
	app := filepath.Base(os.Args[0])
	logfile := filepath.Join(LogPath, app, app+".log")
	Log = logger.New(logger.InfoLevel, logfile, 100, 5)
	log.Println("log file", logfile)
	Log.Infof("logfile: %s", logfile)
	hook.AddCustomStartHook(common.HookEventBeforeShutdown, func() {
		Log.Flush()
	}, common.HookOrderLowest)
}
