package user

import (
	"github.com/gin-gonic/gin"

	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/common/response"
)

func (api *route) profile(c *gin.Context) {
	uid := c.GetInt64("uid")
	profile, err := getInst().userSvc.profile(uid)
	if err != nil {
		c.JSON(httpcode.Http404, response.Error(err))
		return
	}
	if profile == nil {
		c.JSON(httpcode.Http404, response.ErrorNotFound())
		return
	}
	c.JSON(httpcode.Http200, response.Data(profile))
}
