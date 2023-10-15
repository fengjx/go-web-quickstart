package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
	"github.com/fengjx/go-web-quickstart/internal/app/http/router"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/env"
)

type HttpService struct {
	httpServer *http.Server
}

func NewHttpService() *HttpService {
	var serv = &HttpService{}
	return serv
}

func (serv *HttpService) Init() {
	applog.Log.Info("http server init")
	serverConfig := appconfig.Conf.Server
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

	// Register HTTP route handlers.
	router.Init(engine)
	serv.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port),
		Handler: engine,
	}
}

func (serv *HttpService) Start(ctx context.Context) {
	go serv.startHttp(ctx, serv.httpServer)
}

func (serv *HttpService) Shutdown(ctx context.Context) {
	applog.Log.Info("server stop")
	// Graceful HTTP server shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serv.httpServer.Shutdown(ctx); err != nil {
		applog.Log.Error("error while shutting down server: %v", zap.Error(err))
		log.Fatalf("error while shutting down server: %v", err)
	} else {
		log.Println("server was shutdown gracefully")
		applog.Log.Infof("server was shutdown gracefully")
	}
}

// startHttp starts the web server in http mode.
func (serv *HttpService) startHttp(ctx context.Context, srv *http.Server) {
	applog.Log.Infof("http server listening on %s", serv.httpServer.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		applog.Log.Panicf("listen: %s\n", err)
	}
}
