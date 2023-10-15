package app

import (
	"context"

	"github.com/fengjx/go-halo/hook"

	_ "github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/httpclient"
	"github.com/fengjx/go-web-quickstart/internal/app/redis"
	"github.com/fengjx/go-web-quickstart/internal/app/service"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/endpoint"
)

var services []service.Service

func init() {
	db.Init()
	redis.Init()
	httpclient.Init()
	endpoint.Init()
	initServices()
}

func initServices() {
	httpService := service.NewHttpService()
	httpService.Init()
	services = append(services, httpService)
}

func Start(ctx context.Context) {
	hook.DoCustomHooks(common.HookEventBeforeStart)
	for _, svc := range services {
		svc.Start(ctx)
	}
}

func Stop(ctx context.Context) {
	applog.Log.Info("app shutdown")
	// 优先停服务
	for _, svc := range services {
		svc.Shutdown(ctx)
	}
	hook.DoCustomHooks(common.HookEventShutdown)
}
