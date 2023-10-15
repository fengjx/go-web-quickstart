package blog

import (
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/errno"
	"github.com/fengjx/go-web-quickstart/internal/common/response"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
)

type serviceBlog struct {
}

var blogSvc *serviceBlog
var blogSvcInitOnce = sync.Once{}

func getBlogSvc() *serviceBlog {
	blogSvcInitOnce.Do(func() {
		blogSvc = &serviceBlog{}
	})
	return blogSvc
}

func (svc *serviceBlog) Add(blogModel *entity.Blog) (bool, error) {
	blogModel.CreateTime = time.Now().Unix()
	_, err := getBlogDao().Save(blogModel)
	if err != nil {
		applog.Log.Errorf("add blog err - %s", err.Error())
		return false, err
	}
	return true, nil
}

func (svc *serviceBlog) Page(offset int, size int) ([]*entity.Blog, error) {
	return getBlogDao().Page(offset, size)
}

func (svc *serviceBlog) Get(id int64) (*entity.Blog, error) {
	blogModel := &entity.Blog{}
	exist, err := getBlogDao().GetByID(id, blogModel)
	if err != nil {
		applog.Log.Errorf("get blog err - %s", err.Error())
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return blogModel, nil
}

func (svc *serviceBlog) Del(uid int64, id int64) (bool, error) {
	blogModel := &entity.Blog{}
	exist, err := getBlogDao().GetByID(id, blogModel)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, errno.NewErr(response.CodeSystemErr, "user not found")
	}
	if blogModel.Uid != uid {
		return false, errno.NewErr(response.CodeUserErr, "You are not the blog owner")
	}
	ok, err := getBlogDao().DeleteByID(id)
	if err != nil {
		applog.Log.Errorf("del blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}

func (svc *serviceBlog) Update(uid int64, blogModel *entity.Blog) (bool, error) {
	old := &entity.Blog{}
	_, err := getBlogDao().GetByID(blogModel.Id, old)
	if err != nil {
		return false, err
	}
	if old == nil || old.Id == 0 {
		return false, errno.NewErr(response.CodeSystemErr, "blog not exist")
	}
	if old.Uid != uid {
		return false, errno.NewErr(response.CodeUserErr, "You are not the blog owner")
	}
	ok, err := getBlogDao().UpdateField(blogModel.Id, map[string]interface{}{
		"title":   blogModel.Title,
		"content": blogModel.Content,
	})
	if err != nil {
		applog.Log.Errorf("update blog errs", zap.Error(err))
		return false, err
	}
	return ok, nil
}
