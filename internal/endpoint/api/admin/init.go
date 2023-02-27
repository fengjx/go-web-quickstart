package admin

import "github.com/gin-gonic/gin"

func Init(router *gin.Engine) {
	router.Group("/admin")
}
