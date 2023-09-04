package user

import (
	"reflect"
	"sync"

	"github.com/fengjx/daox"

	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/redis"
)

type Dao struct {
	*daox.Dao
}

var once sync.Once
var dao *Dao

func Init() {
	_ = GetDao()
}

func GetDao() *Dao {
	once.Do(func() {
		dao = &Dao{
			Dao: daox.NewDAO(
				db.GetDefaultDB(),
				"user",
				"id",
				reflect.TypeOf(&User{}),
				daox.IsAutoIncrement(),
				daox.WithCache(redis.GetDefaultClient()),
			),
		}
	})
	return dao
}

// GetByUsername
// @description 通过用户名查询用户信息
// @param username 用户名
func (receiver *Dao) GetByUsername(username string) (*User, error) {
	user := &User{}
	exist, err := receiver.GetByColumn(daox.OfKv("username", username), user)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	return user, nil
}
