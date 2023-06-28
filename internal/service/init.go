package service

import "github.com/fengjx/go-web-quickstart/internal/app/applog"

func Init() {
	applog.Log.Infof("service init start")
	_ = GetUserSvc()
	_ = GetBlogSvc()
	applog.Log.Infof("service init end")
}
