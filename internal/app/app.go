package app

import (
	_ "github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/httpclient"
	"github.com/fengjx/go-web-quickstart/internal/app/redis"
	"github.com/fengjx/go-web-quickstart/internal/base/dao"
)

func init() {
	applog.Init()
	db.Init()
	redis.Init()
	httpclient.Init()
	dao.Init()
}

func NewServer() Server {
	var serv Server = &ginServer{}
	return serv
}
