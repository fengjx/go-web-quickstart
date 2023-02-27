package app

import (
	"context"
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
			Log.Error(err)
		}
	}()
	start := time.Now()

	if !IsDev() {
		gin.DisableConsoleColor()
		f, _ := os.Create(Config.Server.Access)
		gin.DefaultWriter = io.MultiWriter(f)
	}
	router := gin.New()
	if IsProd() {
		router.Use(gzip.Gzip(
			gzip.DefaultCompression,
		))
		Log.Infof("server: enabled gzip compression")
	}

	// Register common middleware.
	router.Use(Recovery(), Security())

	// Find and load templates.
	for _, path := range Config.Server.Template {
		router.LoadHTMLFiles(path)
	}

	// Register HTTP route handlers.
	api.RegisterRoutes(router)

	hs := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", Config.Server.Host, Config.Server.Port),
		Handler: router,
	}
	Log.Infof("server: listening on %s [%s]", hs.Addr, time.Since(start))
	go startHttp(hs)

	// Graceful HTTP server shutdown.
	<-ctx.Done()
	Log.Info("server: shutting down")
	err := hs.Close()
	if err != nil {
		Log.Errorf("server: shutdown failed (%s)", err)
	}
}

func (serv *ginServer) Shutdown() {
	Log.Info("server stop")
}

// startHttp starts the web server in http mode.
func startHttp(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			Log.Info("server: shutdown complete")
		} else {
			Log.Errorf("server: %s", err)
		}
	}
}
