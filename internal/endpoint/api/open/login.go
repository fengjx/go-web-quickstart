package open

import (
	"fengjx/go-web-quickstart/internal/app/applog"
	"fengjx/go-web-quickstart/internal/app/http/auth"
	"fengjx/go-web-quickstart/internal/app/http/httpcode"
	"fengjx/go-web-quickstart/internal/common"
	"fengjx/go-web-quickstart/internal/service"
	"github.com/gin-gonic/gin"
)

var userSvc = service.UserService

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
	ok, err := userSvc.Register(req.Username, req.Pwd)
	if err != nil {
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
	u, err := userSvc.Login(req.Username, req.Pwd)
	if err != nil {
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	tokenString, err := auth.Signed(u.Id)
	if err != nil {
		c.JSON(httpcode.Http500, common.UserError("login fail"))
		return
	}
	applog.Log.Infof("user login success: %d", u.Id)
	c.JSON(httpcode.Http200, common.Data(map[string]interface{}{
		"token": tokenString,
	}))
}
