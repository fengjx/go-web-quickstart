package blog

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder"
	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/redis"
	"reflect"
	"sync"
)

type Dao struct {
	*daox.Dao
}

var once sync.Once
var dao *Dao

func GetDao() *Dao {
	once.Do(func() {
		dao = &Dao{
			Dao: daox.NewDAO(
				db.GetDefaultDB(),
				"blog",
				"id",
				reflect.TypeOf(&Blog{}),
				daox.WithCache(redis.GetDefaultClient()),
			),
		}
	})
	return dao
}

func (receiver *Dao) FindUserBlogList(uid int64) ([]*Blog, error) {
	var list []*Blog
	selectSQL, err := receiver.SQLBuilder().Select().Columns(dao.TableMeta.Columns...).Where(
		sqlbuilder.C().Where(true, "uid = ?"),
	).Limit(10).Sql()
	if err != nil {
		return nil, err
	}
	err = dao.DBRead.Select(&list, selectSQL, uid)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (receiver *Dao) Page(offset int, size int) ([]*Blog, error) {
	var list []*Blog
	selectSQL, err := receiver.SQLBuilder().Select().
		Columns(receiver.TableMeta.Columns...).
		Offset(offset).
		Limit(size).
		Sql()
	if err != nil {
		return nil, err
	}
	err = dao.DBRead.Select(&list, selectSQL)
	if err != nil {
		return nil, err
	}
	return list, nil
}
