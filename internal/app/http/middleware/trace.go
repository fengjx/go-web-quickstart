package middleware

import (
	"github.com/fengjx/go-halo/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/fengjx/go-web-quickstart/internal/app/http/header"
)

var Trace = func() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(header.RequestID)
		if reqID == "" {
			reqID = uuid.New().String()
			c.Header(header.RequestID, reqID)
			c.Writer.Header().Set(header.RequestID, reqID)
		}
		c.Set(logger.TraceID, reqID)
	}
}
