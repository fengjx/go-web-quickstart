package service

import (
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/blog"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"time"
)

var BlogService = new(blogService)

type blogService struct {
}

func (receiver *blogService) Add(blogBean *blog.Blog) (bool, error) {
	blogBean.CreateTime = time.Now().UnixMilli()
	ok, err := db.Add(blogBean)
	if err != nil {
		applog.Log.Errorf("add blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}

func (receiver *blogService) Page(offset int, size int) ([]*blog.Blog, error) {
	return blog.Page(offset, size)
}

func (receiver *blogService) Get(id int64) (*blog.Blog, error) {
	blogBean := blog.New()
	err := db.GetById(id, blogBean)
	if err != nil {
		applog.Log.Errorf("get blog err - %s", err.Error())
		return nil, err
	}
	if blogBean.Id == 0 {
		return nil, nil
	}
	return blogBean, nil
}

func (receiver *blogService) Del(uid int64, id int64) (bool, error) {
	blogBean := blog.New()
	err := db.GetById(id, blogBean)
	if err != nil {
		return false, err
	}
	if blogBean.Uid != uid {
		return false, common.NewServiceErr(common.CodeUserErr, "You are not the blog owner")
	}
	ok, err := db.DelById(id, blogBean)
	if err != nil {
		applog.Log.Errorf("del blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}

func (receiver *blogService) Update(uid int64, bean *blog.Blog) (bool, error) {
	old := blog.New()
	err := db.GetById(bean.Id, old)
	if err != nil {
		return false, err
	}
	if old.Uid != uid {
		return false, common.NewServiceErr(common.CodeUserErr, "You are not the blog owner")
	}
	ok, err := db.UpdateById(bean.Id, bean)
	if err != nil {
		applog.Log.Errorf("update blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}
