package health

import (
	"github.com/gin-gonic/gin"

	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/app/http/router"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/response"
)

func initRouter() {
	applog.Log.Info("init health router")
	r := &route{}
	group := router.Root.Group("/health")
	group.GET("/ping", r.ping)
}

type route struct {
}

func (*route) ping(ctx *gin.Context) {
	req := ctx.Request
	data := map[string]interface{}{
		"msg": "pong",
		"data": map[string]interface{}{
			"url":    req.URL,
			"header": req.Header,
		},
	}
	ctx.JSON(httpcode.Http200, response.Data(data))
}
