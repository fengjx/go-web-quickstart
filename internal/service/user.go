package service

import (
	"fengjx/go-web-quickstart/internal/app/applog"
	"fengjx/go-web-quickstart/internal/app/db"
	"fengjx/go-web-quickstart/internal/base/data/user"
	"fengjx/go-web-quickstart/pkg/utils"
	"fmt"
)

var UserService = new(userService)

type userService struct {
}

// Register
// 用户注册
func (svc userService) Register(username string, pwd string) (bool, error) {
	salt := utils.RandomString(6)
	cryptoPwd := utils.Md5SumString(fmt.Sprintf("%s%s", pwd, salt))
	u := user.New()
	u.Username = username
	u.Pwd = cryptoPwd
	u.Salt = salt
	u.Nick = fmt.Sprintf("user-%s", utils.RandomString(8))
	_, err := db.Default().Omit("ctime", "utime").Insert(u)
	if err != nil {
		applog.Log.Errorf("register user err: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (svc userService) Login(username string, pwd string) (*user.User, error) {
	u, err := user.GetByUsername(username)
	if err != nil {
		applog.Log.Errorf("user login err: %s", err.Error())
		return nil, err
	}
	cryptoPwd := utils.Md5SumString(fmt.Sprintf("%s%s", pwd, u.Salt))
	if cryptoPwd != u.Pwd {
		return nil, nil
	}
	return u, nil
}
