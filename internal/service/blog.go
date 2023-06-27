package service

import (
	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/blog"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"time"
)

var BlogService = new(blogService)

type blogService struct {
}

func (receiver *blogService) Add(blogModel *blog.Blog) (bool, error) {
	blogModel.CreateTime = time.Now().UnixMilli()
	_, err := blog.GetDao().Save(blogModel)
	if err != nil {
		applog.Log.Errorf("add blog err - %s", err.Error())
		return false, err
	}
	return true, nil
}

func (receiver *blogService) Page(offset int, size int) ([]*blog.Blog, error) {
	return blog.Page(offset, size)
}

func (receiver *blogService) Get(id int64) (*blog.Blog, error) {
	blogModel := &blog.Blog{}
	err := blog.GetDao().GetByID(id, blogModel)
	if err != nil {
		applog.Log.Errorf("get blog err - %s", err.Error())
		return nil, err
	}
	if blogModel.Id == 0 {
		return nil, nil
	}
	return blogModel, nil
}

func (receiver *blogService) Del(uid int64, id int64) (bool, error) {
	blogModel := &blog.Blog{}
	err := blog.GetDao().GetByID(id, blogModel)
	if err != nil {
		return false, err
	}
	if blogModel.Uid != uid {
		return false, common.NewServiceErr(common.CodeUserErr, "You are not the blog owner")
	}
	ok, err := blog.GetDao().DeleteById(id)
	if err != nil {
		applog.Log.Errorf("del blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}

func (receiver *blogService) Update(uid int64, blogModel *blog.Blog) (bool, error) {
	old := &blog.Blog{}
	err := blog.GetDao().GetByID(blogModel.Id, old)
	if err != nil {
		return false, err
	}
	if old.Uid != uid {
		return false, common.NewServiceErr(common.CodeUserErr, "You are not the blog owner")
	}
	ok, err := blog.GetDao().Update(blogModel)
	if err != nil {
		applog.Log.Errorf("update blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}
