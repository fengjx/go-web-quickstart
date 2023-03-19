package middleware

import (
	"fengjx/go-web-quickstart/internal/app/http/auth"
	"fengjx/go-web-quickstart/internal/app/http/httpcode"
	"fengjx/go-web-quickstart/internal/common"
	"fengjx/go-web-quickstart/internal/common/env"
	"fengjx/go-web-quickstart/pkg/utils"
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
