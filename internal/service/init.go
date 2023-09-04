package service

import (
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/hook"
)

func Init() {
	applog.Log.Infof("service init start")
	initBlogSvc()
	initUserSvc()
	applog.Log.Infof("service init end")
	hook.OnServiceInit()
}
