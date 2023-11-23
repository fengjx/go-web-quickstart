package blog

import (
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/router"
)

func initRouter() {
	applog.Log.Info("init health router")
	r := &route{}
	blog := router.OpenGroup.Group("/blog")
	blog.GET("/", r.index)
	blog.GET("/:id", r.get)
	blog.Use(middleware.Auth()).POST("/", r.add)
	blog.Use(middleware.Auth()).DELETE("/:id", r.del)
	blog.Use(middleware.Auth()).PUT("/:id", r.update)
}

type route struct {
}
