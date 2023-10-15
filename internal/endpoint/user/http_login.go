package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/fengjx/go-web-quickstart/internal/app/http/auth"
	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/response"
)

func (api *route) register(c *gin.Context) {
	req := struct {
		Username string `form:"username" binding:"required"`
		Pwd      string `form:"pwd" binding:"required" min:"6"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, response.ErrorUnauthorized())
		return
	}
	ok, err := getUserSvc().register(req.Username, req.Pwd)
	if err != nil {
		applog.Log.Info("user register err", zap.Error(err))
		c.JSON(httpcode.Http500, response.Error(err))
		return
	}
	c.JSON(httpcode.Http200, response.Status(ok))
}

func (api *route) login(c *gin.Context) {
	req := struct {
		Username string `form:"username" binding:"required"`
		Pwd      string `form:"pwd" binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, response.ErrorBadRequest())
		return
	}
	u, err := getUserSvc().login(req.Username, req.Pwd)
	if err != nil {
		c.JSON(httpcode.Http500, response.Error(err))
		return
	}
	tokenString, err := auth.Signed(u.Id)
	if err != nil {
		c.JSON(httpcode.Http500, response.UserError("login fail"))
		return
	}
	applog.Log.Info("user login success ", zap.Int64("userID", u.Id))
	c.JSON(httpcode.Http200, response.Data(map[string]interface{}{
		"token": tokenString,
	}))
}
