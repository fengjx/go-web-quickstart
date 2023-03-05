package api

import (
	"fengjx/go-web-quickstart/internal/endpoint/api/front"
	"fengjx/go-web-quickstart/internal/endpoint/api/health"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.RedirectTrailingSlash = true
	health.Init(router)
	front.Init(router)
}
