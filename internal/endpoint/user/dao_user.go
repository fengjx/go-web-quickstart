package user

import (
	"reflect"
	"sync"

	"github.com/fengjx/daox"

	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/redis"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
)

type daoUser struct {
	*daox.Dao
}

var once sync.Once
var userDao *daoUser

func getUserDao() *daoUser {
	once.Do(func() {
		userDao = &daoUser{
			Dao: daox.NewDAO(
				db.GetDefaultDB(),
				"user",
				"id",
				reflect.TypeOf(&entity.User{}),
				daox.IsAutoIncrement(),
				daox.WithCache(redis.GetDefaultClient()),
			),
		}
	})
	return userDao
}

// GetByUsername
// @description 通过用户名查询用户信息
// @param username 用户名
func (receiver *daoUser) getByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	exist, err := receiver.GetByColumn(daox.OfKv("username", username), user)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return user, nil
}
