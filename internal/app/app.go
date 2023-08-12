package app

import (
	"context"

	_ "github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/hook"
	"github.com/fengjx/go-web-quickstart/internal/app/httpclient"
	"github.com/fengjx/go-web-quickstart/internal/app/redis"
	"github.com/fengjx/go-web-quickstart/internal/base/dao"
	"github.com/fengjx/go-web-quickstart/internal/service"
)

func init() {
	applog.Init()
	db.Init()
	redis.Init()
	httpclient.Init()
	dao.Init()
	service.Init()
}

func newServer() Server {
	var serv Server = &ginServer{}
	return serv
}

func Start(ctx context.Context) {
	newServer().Start(ctx)
	hook.OnStart()
}

func Stop(ctx context.Context) {
	hook.OnStop()
}
