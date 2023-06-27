package dao

import (
	"github.com/fengjx/go-web-quickstart/internal/base/dao/blog"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/user"
)

func Init() {
	_ = user.GetDao()
	_ = blog.GetDao()
}
