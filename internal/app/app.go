package app

import (
	_ "fengjx/go-web-quickstart/internal/app/appconfig"
	"fengjx/go-web-quickstart/internal/app/applog"
	"fengjx/go-web-quickstart/internal/app/db"
	"fengjx/go-web-quickstart/internal/app/httpclient"
	"fengjx/go-web-quickstart/internal/app/redis"
)

func init() {
	applog.Init()
	db.Init()
	redis.Init()
	httpclient.Init()
}

func NewServer() Server {
	var serv Server = &ginServer{}
	return serv
}
