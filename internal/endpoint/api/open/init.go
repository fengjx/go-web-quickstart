package open

import (
	"github.com/fengjx/go-web-quickstart/internal/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	group := router.Group("/openapi")

	// user
	group.POST("/register", loginApi.register)
	group.POST("/login", loginApi.login)

	// blog
	blog := group.Group("/blog")
	blog.GET("/", blogApi.index)
	blog.GET("/:id", blogApi.get)
	blog.Use(middleware.Auth()).POST("/", blogApi.add)
	blog.Use(middleware.Auth()).DELETE("/:id", blogApi.del)
	blog.Use(middleware.Auth()).PUT("/:id", blogApi.update)

	user := group.Group("/user")
	user.Use(middleware.Auth()).GET("/profile", userApi.profile)

}
