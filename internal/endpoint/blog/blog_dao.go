package blog

import (
	"reflect"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder/ql"

	"github.com/fengjx/go-web-quickstart/internal/app/db"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
	"github.com/fengjx/go-web-quickstart/internal/data/meta"
)

type _blogDao struct {
	*daox.Dao
	m meta.BlogM
}

func newBlogDao() *_blogDao {
	blogDao := &_blogDao{
		Dao: daox.NewDAO(
			db.GetDefaultDB(),
			"blog",
			"id",
			reflect.TypeOf(&entity.Blog{}),
		),
		m: meta.BlogMeta,
	}
	return blogDao
}

func (dao *_blogDao) findUserBlogList(uid int64) ([]*entity.Blog, error) {
	var list []*entity.Blog
	selector := dao.Selector().Where(ql.EC().Where(dao.m.UidEQ(uid))).Limit(10)
	err := dao.Select(&list, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (dao *_blogDao) page(offset int, size int) ([]*entity.Blog, error) {
	var list []*entity.Blog
	selector := dao.SQLBuilder().Select().
		Offset(offset).
		Limit(size)
	err := dao.Select(&list, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}
