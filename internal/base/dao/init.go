package dao

import (
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/blog"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/user"
)

func Init() {
	applog.Log.Infof("dao init start")
	_ = user.GetDao()
	_ = blog.GetDao()
	applog.Log.Infof("dao init end")
}
