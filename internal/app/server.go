package app

import (
	"context"
	"fengjx/go-web-quickstart/internal/app/appconfig"
	"fengjx/go-web-quickstart/internal/app/applog"
	"fengjx/go-web-quickstart/internal/common/env"
	"fengjx/go-web-quickstart/internal/endpoint/api"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start(ctx context.Context)
	Shutdown()
}

type ginServer struct {
}

func (serv *ginServer) Start(ctx context.Context) {
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
	if env.IsProd() {
		gin.SetMode(gin.ReleaseMode)
		router.Use(gzip.Gzip(
			gzip.DefaultCompression,
		))
		applog.Log.Infof("server: enabled gzip compression")
	}

	// Register common middleware.
	router.Use(Recovery(), Security())

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
	applog.Log.Infof("server: listening on %s [%s]", hs.Addr, time.Since(start))
	go startHttp(hs)

	// Graceful HTTP server shutdown.
	<-ctx.Done()
	applog.Log.Info("server: shutting down")
	err := hs.Close()
	if err != nil {
		applog.Log.Errorf("server: shutdown failed (%s)", err)
	}
}

func (serv *ginServer) Shutdown() {
	applog.Log.Info("server stop")
}

// startHttp starts the web server in http mode.
func startHttp(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			applog.Log.Info("server: shutdown complete")
		} else {
			applog.Log.Errorf("server: %s", err)
		}
	}
}
