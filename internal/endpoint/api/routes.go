package api

import (
	"fengjx/go-web-quickstart/internal/endpoint/api/health"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes configures the available web server routes.
func RegisterRoutes(router *gin.Engine) {
	// Enables automatic redirection if the current route cannot be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	health.Init(router)

}
