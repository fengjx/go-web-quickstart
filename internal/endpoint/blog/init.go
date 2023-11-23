package blog

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
	blogDao *_blogDao
	blogSvc *_blogService
}

var ins *inst
var insOnce sync.Once

func getInst() *inst {
	insOnce.Do(func() {
		ins = &inst{
			blogDao: newBlogDao(),
			blogSvc: newBlogSvc(),
		}
	})
	return ins
}
