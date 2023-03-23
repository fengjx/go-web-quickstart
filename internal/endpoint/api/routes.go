package api

import (
	"github.com/fengjx/go-web-quickstart/internal/endpoint/api/admin"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/api/health"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/api/open"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.RedirectTrailingSlash = true
	health.Init(router)
	open.Init(router)
	admin.Init(router)
}
