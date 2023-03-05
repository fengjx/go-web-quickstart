package blog

import "fengjx/go-web-quickstart/internal/app/db"

func FindUserBlogList(uid int64) ([]*Blog, error) {
	list := make([]*Blog, 0)
	err := db.Default().Where("uid = ?", uid).Limit(10).Find(list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
