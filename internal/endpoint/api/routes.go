package api

import (
	"fengjx/go-web-quickstart/internal/endpoint/api/health"
	"fengjx/go-web-quickstart/internal/endpoint/api/open"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.RedirectTrailingSlash = true
	health.Init(router)
	open.Init(router)
}
