package open

import (
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/http/httpcode"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/blog"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	blogService = service.BlogService
)

var blogApi *BlogApi

func init() {
	blogApi = new(BlogApi)
}

type BlogApi struct {
}

func (api *BlogApi) index(c *gin.Context) {
	req := struct {
		Offset int
		Size   int
	}{
		Offset: utils.ToInt(c.DefaultQuery("offset", "0")),
		Size:   utils.ToInt(c.DefaultQuery("size", "10")),
	}
	list, err := blogService.Page(req.Offset, req.Size)
	if err != nil {
		applog.Log.Errorf("query page err - %+v", err)
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	c.JSON(httpcode.Http200, common.Data(list))
}

func (api *BlogApi) get(c *gin.Context) {
	req := struct {
		Id int64 `binding:"required"`
	}{
		Id: utils.ToInt64(c.Param("id")),
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http500, common.UserError(err.Error()))
		return
	}
	blogBean, err := blogService.Get(req.Id)
	if err != nil {
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	c.JSON(httpcode.Http200, common.Data(blogBean))
}

func (api *BlogApi) add(c *gin.Context) {
	uid := c.GetInt64("uid")
	req := struct {
		Title   string `binding:"required"`
		Content string `binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, common.UserError(err.Error()))
		return
	}
	ok, err := blogService.Add(&blog.Blog{
		Uid:     uid,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	c.JSON(httpcode.Http200, common.Status(ok))
}

func (api *BlogApi) update(c *gin.Context) {
	uid := c.GetInt64("uid")
	req := struct {
		Id      int64 `binding:"required"`
		Title   string
		Content string
	}{
		Id: utils.ToInt64(c.Param("id")),
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, common.UserError(err.Error()))
		return
	}
	if req.Title == "" && req.Content == "" {
		c.JSON(httpcode.Http400, common.NewServiceErr(common.CodeErrBadRequest, "title and content all nil"))
		return
	}
	bean := &blog.Blog{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
	}
	_, err := blogService.Update(uid, bean)
	if err != nil {
		applog.Log.Errorf("update blog err - %s", err.Error())
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	c.JSON(httpcode.Http200, common.OK())
}

func (api *BlogApi) del(c *gin.Context) {
	uid := c.GetInt64("uid")
	req := struct {
		Id int64 `binding:"required"`
	}{
		Id: utils.ToInt64(c.Param("id")),
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(httpcode.Http400, common.UserError(err.Error()))
		return
	}
	ok, err := blogService.Del(uid, req.Id)
	if err != nil {
		c.JSON(httpcode.Http500, common.Error(err))
		return
	}
	c.JSON(httpcode.Http200, common.Status(ok))
}
