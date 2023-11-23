package user

import (
	"reflect"

	"github.com/fengjx/daox"

	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
	"github.com/fengjx/go-web-quickstart/internal/data/meta"
)

type _userDao struct {
	*daox.Dao
	m meta.UserM
}

func newUserDao() *_userDao {
	userDao := &_userDao{
		Dao: daox.NewDAO(
			db.GetDefaultDB(),
			"user",
			"id",
			reflect.TypeOf(&entity.User{}),
			daox.IsAutoIncrement(),
		),
		m: meta.UserMeta,
	}
	return userDao
}

// GetByUsername
// @description 通过用户名查询用户信息
// @param username 用户名
func (dao *_userDao) getByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	exist, err := dao.GetByColumn(daox.OfKv("username", username), user)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return user, nil
}
