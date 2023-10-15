package blog

import (
	"github.com/fengjx/go-halo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/response"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
)

func (api *route) index(c *gin.Context) {
	req := struct {
		Offset int
		Size   int
	}{
		Offset: utils.ToInt(c.DefaultQuery("offset", "0")),
		Size:   utils.ToInt(c.DefaultQuery("size", "10")),
	}
	list, err := getBlogSvc().Page(req.Offset, req.Size)
	if err != nil {
		applog.Log.Errorf("query page err", zap.Error(err))
		c.JSON(httpcode.Http500, response.Error(err))
		return
	}
	c.JSON(httpcode.Http200, response.Data(list))
}

func (api *route) get(c *gin.Context) {
	req := struct {
		Id int64 `binding:"required"`
	}{
		Id: utils.ToInt64(c.Param("id")),
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http500, response.UserError(err.Error()))
		return
	}
	blogBean, err := getBlogSvc().Get(req.Id)
	if err != nil {
		c.JSON(httpcode.Http500, response.Error(err))
		return
	}
	c.JSON(httpcode.Http200, response.Data(blogBean))
}

func (api *route) add(c *gin.Context) {
	uid := c.GetInt64("uid")
	req := struct {
		Title   string `binding:"required"`
		Content string `binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, response.UserError(err.Error()))
		return
	}
	ok, err := getBlogSvc().Add(&entity.Blog{
		Uid:     uid,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		c.JSON(httpcode.Http500, response.Error(err))
		return
	}
	c.JSON(httpcode.Http200, response.Status(ok))
}

func (api *route) update(c *gin.Context) {
	uid := c.GetInt64("uid")
	req := struct {
		Id      int64 `binding:"required"`
		Title   string
		Content string
	}{
		Id: utils.ToInt64(c.Param("id")),
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, response.UserError(err.Error()))
		return
	}
	if req.Title == "" && req.Content == "" {
		c.JSON(httpcode.Http400, response.ErrorBadRequestWithMsg("title and content all nil"))
		return
	}
	bean := &entity.Blog{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	}
	_, err := getBlogSvc().Update(uid, bean)
	if err != nil {
		applog.Log.Errorf("update blog err - %s", err.Error())
		c.JSON(httpcode.Http500, response.Error(err))
		return
	}
	c.JSON(httpcode.Http200, response.OK())
}

func (api *route) del(c *gin.Context) {
	uid := c.GetInt64("uid")
	req := struct {
		Id int64 `binding:"required"`
	}{
		Id: utils.ToInt64(c.Param("id")),
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, response.UserError(err.Error()))
		return
	}
	ok, err := getBlogSvc().Del(uid, req.Id)
	if err != nil {
		c.JSON(httpcode.Http500, response.Error(err))
		return
	}
	c.JSON(httpcode.Http200, response.Status(ok))
}
