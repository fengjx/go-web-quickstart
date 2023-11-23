package app

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/fengjx/go-halo/hook"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
	"github.com/fengjx/go-web-quickstart/internal/app/service"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/env"
	"github.com/fengjx/go-web-quickstart/internal/endpoint"
	"github.com/fengjx/go-web-quickstart/internal/facade"
)

var services []service.Service

func createEngine() *gin.Engine {
	serverConfig := appconfig.GetConfig().Server
	if env.IsTest() || env.IsProd() {
		gin.DisableConsoleColor()
		f, err := os.Create(filepath.Join(applog.LogPath, "access.log"))
		if err != nil {
			panic(err)
		}
		gin.DefaultWriter = io.MultiWriter(f)
	}
	engine := gin.New()
	pprof.Register(engine)
	if env.IsProd() {
		gin.SetMode(gin.ReleaseMode)
		engine.Use(gzip.Gzip(
			gzip.DefaultCompression,
		))
		applog.Log.Infof("server: enabled gzip compression")
	}
	// Register common middleware.
	engine.Use(
		middleware.Trace(),
		middleware.Recovery(),
		middleware.Security(),
	)
	// Find and load templates.
	if len(serverConfig.Template) > 0 {
		engine.LoadHTMLFiles(serverConfig.Template...)
	}
	return engine
}

func initServices(engine *gin.Engine) {
	httpService := service.NewHttpService(engine)
	services = append(services, httpService)
}

func startService(ctx context.Context) {
	hook.DoCustomHooks(common.HookEventBeforeStart)
	for _, svc := range services {
		svc.Start(ctx)
	}
	hook.DoCustomHooks(common.HookEventAfterStart)
}

func Start(ctx context.Context) {
	engine := createEngine()
	db.Init()
	facade.Init()
	endpoint.Init(engine)
	initServices(engine)
	startService(ctx)
}

func Stop(ctx context.Context) {
	// 优先停服务
	hook.DoCustomHooks(common.HookEventBeforeShutdown)
	for _, svc := range services {
		svc.Shutdown(ctx)
	}
	hook.DoCustomHooks(common.HookEventAfterShutdown)
}
