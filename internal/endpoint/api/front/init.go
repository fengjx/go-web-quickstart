package front

import "github.com/gin-gonic/gin"

var group *gin.RouterGroup

func Init(router *gin.Engine) {
	group = router.Group("/api")
	group.POST("/register", register)
	group.POST("/login", login)
}
