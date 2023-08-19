package service

import (
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/hook"
)

func Init() {
	_ = GetUserSvc()
	_ = GetBlogSvc()
	applog.Log.Infof("service init")
	hook.OnServiceInit()
}
