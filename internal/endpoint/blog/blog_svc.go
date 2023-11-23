package blog

import (
	"time"

	"go.uber.org/zap"

	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/errno"
	"github.com/fengjx/go-web-quickstart/internal/common/response"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
)

type _blogService struct {
}

func newBlogSvc() *_blogService {
	return &_blogService{}
}

func (svc *_blogService) add(blogModel *entity.Blog) (bool, error) {
	blogModel.CreateTime = time.Now().Unix()
	_, err := getInst().blogDao.Save(blogModel)
	if err != nil {
		applog.Log.Errorf("add blog err - %s", err.Error())
		return false, err
	}
	return true, nil
}

func (svc *_blogService) page(offset int, size int) ([]*entity.Blog, error) {
	return getInst().blogDao.page(offset, size)
}

func (svc *_blogService) get(id int64) (*entity.Blog, error) {
	blogModel := &entity.Blog{}
	exist, err := getInst().blogDao.GetByID(id, blogModel)
	if err != nil {
		applog.Log.Errorf("get blog err - %s", err.Error())
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return blogModel, nil
}

func (svc *_blogService) del(uid int64, id int64) (bool, error) {
	blogModel := &entity.Blog{}
	exist, err := getInst().blogDao.GetByID(id, blogModel)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, errno.NewErr(response.CodeSystemErr, "user not found")
	}
	if blogModel.UID != uid {
		return false, errno.NewErr(response.CodeUserErr, "You are not the blog owner")
	}
	ok, err := getInst().blogDao.DeleteByID(id)
	if err != nil {
		applog.Log.Errorf("del blog err - %s", err.Error())
		return false, err
	}
	return ok, nil
}

func (svc *_blogService) update(uid int64, blogModel *entity.Blog) (bool, error) {
	old := &entity.Blog{}
	_, err := getInst().blogDao.GetByID(blogModel.ID, old)
	if err != nil {
		return false, err
	}
	if old == nil || old.ID == 0 {
		return false, errno.NewErr(response.CodeSystemErr, "blog not exist")
	}
	if old.UID != uid {
		return false, errno.NewErr(response.CodeUserErr, "You are not the blog owner")
	}
	ok, err := getInst().blogDao.UpdateField(blogModel.ID, map[string]interface{}{
		"title":   blogModel.Title,
		"content": blogModel.Content,
	})
	if err != nil {
		applog.Log.Errorf("update blog errs", zap.Error(err))
		return false, err
	}
	return ok, nil
}
