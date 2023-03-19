package blog

import "fengjx/go-web-quickstart/internal/app/db"

func FindUserBlogList(uid int64) ([]*Blog, error) {
	var list []*Blog
	err := db.Default().Where("uid = ?", uid).Limit(10).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Page(offset int, size int) ([]*Blog, error) {
	var list []*Blog
	err := db.Default().Limit(size, offset).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
