package blog

import (
	"reflect"
	"sync"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder"

	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/app/redis"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
)

type daoBlog struct {
	*daox.Dao
}

var blogDao *daoBlog
var blogDaoOnce sync.Once

func getBlogDao() *daoBlog {
	blogDaoOnce.Do(func() {
		blogDao = &daoBlog{
			Dao: daox.NewDAO(
				db.GetDefaultDB(),
				"blog",
				"id",
				reflect.TypeOf(&entity.Blog{}),
				daox.WithCache(redis.GetDefaultClient()),
			),
		}
	})
	return blogDao
}

func (dao *daoBlog) FindUserBlogList(uid int64) ([]*entity.Blog, error) {
	var list []*entity.Blog
	selectSQL, err := dao.SQLBuilder().Select().Columns(dao.TableMeta.Columns...).Where(
		sqlbuilder.C().Where(true, "uid = ?"),
	).Limit(10).SQL()
	if err != nil {
		return nil, err
	}
	err = dao.DBRead.Select(&list, selectSQL, uid)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (dao *daoBlog) Page(offset int, size int) ([]*entity.Blog, error) {
	var list []*entity.Blog
	selectSQL, err := dao.SQLBuilder().Select().
		Columns(dao.TableMeta.Columns...).
		Offset(offset).
		Limit(size).
		SQL()
	if err != nil {
		return nil, err
	}
	err = dao.DBRead.Select(&list, selectSQL)
	if err != nil {
		return nil, err
	}
	return list, nil
}
