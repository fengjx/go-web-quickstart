package user

import "fengjx/go-web-quickstart/internal/app/db"

// GetByUsername
// @description 通过用户名查询用户信息
// @param username 用户名
func GetByUsername(username string) (*User, error) {
	user := New()
	_, err := db.Default().Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
