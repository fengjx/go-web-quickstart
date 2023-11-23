package user

import (
	"sync"

	"github.com/fengjx/go-halo/hook"

	"github.com/fengjx/go-web-quickstart/internal/common"
)

func Init() {
	hook.AddCustomStartHook(common.HookEventBeforeStart, func() {
		initRouter()
	}, common.HookOrderNormal)
}

type inst struct {
	userDao *_userDao
	userSvc *_userService
}

var ins *inst
var insOnce sync.Once

func getInst() *inst {
	insOnce.Do(func() {
		ins = &inst{
			userDao: newUserDao(),
			userSvc: newUserSvc(),
		}
	})
	return ins
}
