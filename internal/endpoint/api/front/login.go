package front

import (
	"fengjx/go-web-quickstart/internal/app/http/response"
	"fengjx/go-web-quickstart/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userSvc = service.UserService

type RegisterParam struct {
	Username string `form:"username" binding:"required"`
	Pwd      string `form:"pwd" binding:"required" min:"6"`
}

func register(ctx *gin.Context) {
	var param RegisterParam
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, response.UserError(err.Error()))
		return
	}
	ok, err := userSvc.Register(param.Username, param.Pwd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.UserError("system error"))
		return
	}
	if !ok {
		ctx.JSON(http.StatusInternalServerError, response.UserError("register fail, please retry"))
		return
	}
	ctx.JSON(http.StatusOK, response.OK())
}

type LoginParam struct {
	Username string `form:"username" binding:"required"`
	Pwd      string `form:"pwd" binding:"required"`
}

func login(ctx *gin.Context) {
	var param LoginParam
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, response.UserError(err.Error()))
		return
	}
	u, err := userSvc.Login(param.Username, param.Pwd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.UserError("system error"))
		return
	}
	if u == nil {
		ctx.JSON(http.StatusInternalServerError, response.UserError("username or password error"))
		return
	}

}
