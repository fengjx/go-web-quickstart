package endpoint

import (
	"github.com/fengjx/go-web-quickstart/internal/endpoint/blog"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/health"
	"github.com/fengjx/go-web-quickstart/internal/endpoint/user"
)

func Init() {
	health.Init()
	user.Init()
	blog.Init()
}
