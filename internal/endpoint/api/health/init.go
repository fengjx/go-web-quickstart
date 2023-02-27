package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	group := router.Group("/health")
	group.GET("/ping", ping)
}

func ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
