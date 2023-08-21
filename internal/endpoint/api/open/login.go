package open

import (
	"github.com/gin-gonic/gin"

	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/http/auth"
	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/service"
)

var loginApi *LoginApi

func init() {
	loginApi = new(LoginApi)
}

type LoginApi struct {
}

func (api *LoginApi) register(c *gin.Context) {
	req := struct {
		Username string `form:"username" binding:"required"`
		Pwd      string `form:"pwd" binding:"required" min:"6"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, common.ErrorUnauthorized())
		return
	}
	ok, err := service.GetUserSvc().Register(req.Username, req.Pwd)
	if err != nil {
		applog.Log.Info("user register err - ", err)
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	c.JSON(httpcode.Http200, common.Status(ok))
}

func (api *LoginApi) login(c *gin.Context) {
	req := struct {
		Username string `form:"username" binding:"required"`
		Pwd      string `form:"pwd" binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, common.ErrorBadRequest())
		return
	}
	u, err := service.GetUserSvc().Login(req.Username, req.Pwd)
	if err != nil {
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	tokenString, err := auth.Signed(u.Id)
	if err != nil {
		c.JSON(httpcode.Http500, common.UserError("login fail"))
		return
	}
	applog.Log.With(c.Copy()).Infof("user login success: %d", u.Id)
	c.JSON(httpcode.Http200, common.Data(map[string]interface{}{
		"token": tokenString,
	}))
}
