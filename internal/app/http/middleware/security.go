package middleware

import (
	"github.com/fengjx/go-web-quickstart/internal/app/http/header"
	"github.com/gin-gonic/gin"
)

// Security adds common HTTP security headers to the response.
var Security = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(header.ContentSecurityPolicy, header.DefaultContentSecurityPolicy)
		c.Writer.Header().Set(header.FrameOptions, header.DefaultFrameOptions)
	}
}
