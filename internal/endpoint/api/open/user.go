package open

import (
	"fengjx/go-web-quickstart/internal/app/http/httpcode"
	"fengjx/go-web-quickstart/internal/common"
	"fengjx/go-web-quickstart/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	userService = service.UserService
)

var userApi = new(UserApi)

type UserApi struct {
}

func (api *UserApi) profile(c *gin.Context) {
	uid := c.GetInt64("uid")
	profile, err := userService.Profile(uid)
	if err != nil {
		c.JSON(httpcode.Http404, common.Error(err))
		return
	}
	if profile == nil {
		c.JSON(httpcode.Http404, common.ErrorNotFound())
		return
	}
	c.JSON(httpcode.Http200, common.Data(profile))
}
