package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
)

var Root *gin.Engine
var AdminGroup *gin.RouterGroup
var OpenGroup *gin.RouterGroup

func Init(engine *gin.Engine) {
	Root = engine
	if appconfig.GetConfig().Server.Static != "" {
		Root.GET("/", index)
		Root.Static("ui", appconfig.GetConfig().Server.Static)
	}
	AdminGroup = Root.Group("/admin")
	AdminGroup.Use(middleware.Admin())

	OpenGroup = Root.Group("/openapi")
}

func index(c *gin.Context) {
	c.Redirect(http.StatusFound, "/ui")
}
