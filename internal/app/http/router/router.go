package router

import (
	"github.com/gin-gonic/gin"

	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
)

var Root *gin.Engine
var AdminGroup *gin.RouterGroup
var OpenGroup *gin.RouterGroup

func Init(engine *gin.Engine) {
	Root = engine
	AdminGroup = Root.Group("/admin")
	AdminGroup.Use(middleware.Admin())

	OpenGroup = Root.Group("/openapi")
}

func IsEngineInit() bool {
	return Root != nil
}
