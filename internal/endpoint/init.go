package endpoint

import (
	"github.com/gin-gonic/gin"

	"github.com/fengjx/go-web-quickstart/internal/endpoint/blog"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/health"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/router"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/user"
)

func Init(engine *gin.Engine) {
	router.Init(engine)
	health.Init()
	user.Init()
	blog.Init()
}
