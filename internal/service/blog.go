package service

import (
	"sync"
	"time"

	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/blog"
	"github.com/fengjx/go-web-quickstart/internal/common"
)

type BlogService struct {
	blogDao *blog.Dao
}

var blogSvc *BlogService
var blogSvcInitOnce = sync.Once{}

func GetBlogSvc() *BlogService {
	blogSvcInitOnce.Do(func() {
		blogSvc = &BlogService{
			blogDao: blog.GetDao(),
		}
	})
	return blogSvc
}

func (receiver *BlogService) Add(blogModel *blog.Blog) (bool, error) {
	blogModel.CreateTime = time.Now().Unix()
	_, err := receiver.blogDao.Save(blogModel)
	if err != nil {
		applog.Log.Errorf("add blog err - %s", err.Error())
		return false, err
	}
	return true, nil
}

func (receiver *BlogService) Page(offset int, size int) ([]*blog.Blog, error) {
	return receiver.blogDao.Page(offset, size)
}

func (receiver *BlogService) Get(id int64) (*blog.Blog, error) {
	blogModel := &blog.Blog{}
	err := receiver.blogDao.GetByID(id, blogModel)
	if err != nil {
		applog.Log.Errorf("get blog err - %s", err.Error())
		return nil, err
	}
	if blogModel.Id == 0 {
		return nil, nil
	}
	return blogModel, nil
}

func (receiver *BlogService) Del(uid int64, id int64) (bool, error) {
	blogModel := &blog.Blog{}
	err := receiver.blogDao.GetByID(id, blogModel)
	if err != nil {
		return false, err
	}
	if blogModel.Uid != uid {
		return false, common.NewServiceErr(common.CodeUserErr, "You are not the blog owner")
	}
	ok, err := receiver.blogDao.DeleteById(id)
	if err != nil {
		applog.Log.Errorf("del blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}

func (receiver *BlogService) Update(uid int64, blogModel *blog.Blog) (bool, error) {
	old := &blog.Blog{}
	err := receiver.blogDao.GetByID(blogModel.Id, old)
	if err != nil {
		return false, err
	}
	if old == nil || old.Id == 0 {
		return false, common.NewServiceErr(common.CodeUserErr, "blog not exist")
	}
	if old.Uid != uid {
		return false, common.NewServiceErr(common.CodeUserErr, "You are not the blog owner")
	}
	ok, err := receiver.blogDao.UpdateField(blogModel.Id, map[string]interface{}{
		"title":   blogModel.Title,
		"content": blogModel.Content,
	})
	if err != nil {
		applog.Log.Errorf("update blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}
