package user

import (
	"fmt"
	"sync"

	"github.com/fengjx/go-halo/utils"
	"go.uber.org/zap"

	"github.com/fengjx/go-web-quickstart/internal/common/applog"
	"github.com/fengjx/go-web-quickstart/internal/common/errno"
	"github.com/fengjx/go-web-quickstart/internal/common/response"
	"github.com/fengjx/go-web-quickstart/internal/data/dto"
	"github.com/fengjx/go-web-quickstart/internal/data/entity"
)

type serviceUser struct {
}

var userSvc *serviceUser
var userSvcOnce = sync.Once{}

func getUserSvc() *serviceUser {
	userSvcOnce.Do(func() {
		userSvc = &serviceUser{}
	})
	return userSvc
}

// Register
// 用户注册
func (svc *serviceUser) register(username string, pwd string) (bool, error) {
	old, err := getUserDao().getByUsername(username)
	if err != nil {
		return false, err
	}
	if old != nil && old.Username == username {
		return false, errno.NewErr(response.CodeUserErr, "username already exists")
	}
	salt := utils.RandomString(6)
	cryptoPwd := utils.Md5SumString(fmt.Sprintf("%s%s", pwd, salt))
	u := &entity.User{
		Username: username,
		Nick:     fmt.Sprintf("user-%s", utils.RandomString(8)),
		Pwd:      cryptoPwd,
		Salt:     salt,
	}
	_, err = getUserDao().Save(u, "ctime", "utime")
	if err != nil {
		applog.Log.Errorf("register user err: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (svc *serviceUser) login(username string, pwd string) (*entity.User, error) {
	u, err := getUserDao().getByUsername(username)
	if err != nil {
		applog.Log.Error("user login err", zap.Error(err))
		return nil, err
	}
	if u.Id == 0 {
		return nil, errno.NewErr(response.CodeUserErr, "username not exists")
	}
	cryptoPwd := utils.Md5SumString(fmt.Sprintf("%s%s", pwd, u.Salt))
	if cryptoPwd != u.Pwd {
		return nil, errno.NewErr(response.CodeUserErr, "password invalid")
	}
	return u, nil
}

func (svc *serviceUser) profile(uid int64) (*dto.UserDTO, error) {
	u := &entity.User{}
	exist, err := getUserDao().GetByID(uid, u)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	userDTO := &dto.UserDTO{}
	userDTO.Of(u)
	return userDTO, nil
}
