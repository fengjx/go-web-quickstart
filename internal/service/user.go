package service

import (
	"fmt"
	"sync"

	"github.com/fengjx/go-halo/utils"

	"github.com/fengjx/go-web-quickstart/internal/app/applog"
	"github.com/fengjx/go-web-quickstart/internal/base/dao/user"
	"github.com/fengjx/go-web-quickstart/internal/common"
	"github.com/fengjx/go-web-quickstart/internal/dto"
)

type UserService struct {
	userDao *user.Dao
}

var userSvc = new(UserService)
var userSvcInitOnce = sync.Once{}

func GetUserSvc() *UserService {
	userSvcInitOnce.Do(func() {
		userSvc = &UserService{
			userDao: user.GetDao(),
		}
	})
	return userSvc
}

// Register
// 用户注册
func (receiver *UserService) Register(username string, pwd string) (bool, error) {
	old, err := receiver.userDao.GetByUsername(username)
	if err != nil {
		return false, err
	}
	if old.Username == username {
		return false, common.NewServiceErr(common.CodeUserErr, "username already exists")
	}
	salt := utils.RandomString(6)
	cryptoPwd := utils.Md5SumString(fmt.Sprintf("%s%s", pwd, salt))
	u := &user.User{
		Username: username,
		Nick:     fmt.Sprintf("user-%s", utils.RandomString(8)),
		Pwd:      cryptoPwd,
		Salt:     salt,
	}
	_, err = user.GetDao().Save(u, "ctime", "utime")
	if err != nil {
		applog.Log.Errorf("register user err: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (receiver *UserService) Login(username string, pwd string) (*user.User, error) {
	u, err := receiver.userDao.GetByUsername(username)
	if err != nil {
		applog.Log.With(nil).Errorf("user login err: %s", err.Error())
		return nil, err
	}
	if u.Id == 0 {
		return nil, common.NewServiceErr(common.CodeUserErr, "username not exists")
	}
	cryptoPwd := utils.Md5SumString(fmt.Sprintf("%s%s", pwd, u.Salt))
	if cryptoPwd != u.Pwd {
		return nil, common.NewServiceErr(common.CodeUserErr, "password invalid")
	}
	return u, nil
}

func (receiver *UserService) Profile(uid int64) (*dto.UserDTO, error) {
	u := &user.User{}
	err := user.GetDao().GetByID(uid, u)
	if err != nil {
		return nil, err
	}
	if u.Id == 0 {
		return nil, nil
	}
	userDTO := &dto.UserDTO{}
	userDTO.Of(u)
	return userDTO, nil
}
