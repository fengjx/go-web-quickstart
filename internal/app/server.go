package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/pprof"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/hook"
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
	"github.com/fengjx/go-web-quickstart/internal/common/env"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/api"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start()
	Shutdown()
}

type ginServer struct {
}

func (serv *ginServer) Start() {
	defer func() {
		if err := recover(); err != nil {
			applog.Log.Error(err)
		}
	}()
	start := time.Now()

	serverConfig := appconfig.Conf.Server

	if !env.IsDev() {
		gin.DisableConsoleColor()
		f, _ := os.Create(serverConfig.Access)
		gin.DefaultWriter = io.MultiWriter(f)
	}
	router := gin.New()
	pprof.Register(router)
	if env.IsProd() {
		gin.SetMode(gin.ReleaseMode)
		router.Use(gzip.Gzip(
			gzip.DefaultCompression,
		))
		applog.Log.Infof("server: enabled gzip compression")
	}

	// Register common middleware.
	router.Use(
		middleware.Recovery(),
		middleware.Security(),
	)

	// Find and load templates.
	for _, path := range serverConfig.Template {
		router.LoadHTMLFiles(path)
	}

	// Register HTTP route handlers.
	api.RegisterRoutes(router)

	hs := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port),
		Handler: router,
	}
	applog.Log.Infof("server listening on %s [%s]", hs.Addr, time.Since(start))
	go startHttp(hs)
}

func (serv *ginServer) Shutdown() {
	applog.Log.Info("server stop")
}

// startHttp starts the web server in http mode.
func startHttp(srv *http.Server) {
	// 首先退出 http server 停止用户请求
	hook.AddStopHook(func() {
		// Graceful HTTP server shutdown.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			applog.Log.Error("error while shutting down server: %v", err)
			log.Fatalf("error while shutting down server: %v", err)
		} else {
			log.Println("server was shutdown gracefully")
			applog.Log.Infof("server was shutdown gracefully")
		}
	}, 1)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
