package service

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
)

type HttpService struct {
	httpServer *http.Server
}

func NewHttpService(engine *gin.Engine) *HttpService {
	serverConfig := appconfig.GetConfig().Server
	httpServer := &http.Server{
		Addr:    serverConfig.Listen,
		Handler: engine,
	}
	var serv = &HttpService{
		httpServer: httpServer,
	}
	return serv
}

func (serv *HttpService) Start(ctx context.Context) {
	go serv.startHttp(ctx, serv.httpServer)
}

func (serv *HttpService) Shutdown(_ context.Context) {
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
func (serv *HttpService) startHttp(_ context.Context, srv *http.Server) {
	applog.Log.Infof("server listening on %s", serv.httpServer.Addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
