package user

import (
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/router"
)

func initRouter() {
	applog.Log.Info("init user router")
	r := &route{}
	// user login
	router.OpenGroup.POST("/register", r.register)
	router.OpenGroup.POST("/login", r.login)

	user := router.OpenGroup.Group("/user")
	user.Use(middleware.Auth()).GET("/profile", r.profile)
}

type route struct {
}
