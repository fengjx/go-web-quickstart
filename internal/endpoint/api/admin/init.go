package admin

import (
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
	"github.com/gin-gonic/gin"
)

var group *gin.RouterGroup

func Init(router *gin.Engine) {
	group = router.Group("/admin")
	group.Use(middleware.Admin())
}
