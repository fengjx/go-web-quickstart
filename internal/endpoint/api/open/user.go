package open

import (
	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/service"
	"github.com/gin-gonic/gin"
)

var userApi = new(UserApi)

type UserApi struct {
}

func (api *UserApi) profile(c *gin.Context) {
	uid := c.GetInt64("uid")
	profile, err := service.GetUserSvc().Profile(uid)
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
