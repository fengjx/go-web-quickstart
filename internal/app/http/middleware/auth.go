package middleware

import (
	"github.com/fengjx/go-web-quickstart/internal/app/http/auth"
	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/common/env"
	"github.com/fengjx/go-web-quickstart/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uid int64
		if !env.IsProd() {
			uid = utils.ToInt64(ctx.GetHeader("X-Uid"))
		}
		if uid > 0 {
			ctx.Set("uid", uid)
			ctx.Next()
			return
		}
		token := ctx.GetHeader("X-Token")
		if token == "" {
			ctx.AbortWithStatusJSON(httpcode.Http401, common.ErrorUnauthorized())
			return
		}
		uid, err := auth.Parse(token)
		if uid == 0 || err != nil {
			ctx.AbortWithStatusJSON(httpcode.Http401, common.ErrorUnauthorized())
			return
		}
		ctx.Set("uid", uid)
	}
}
