package health

import (
	"github.com/fengjx/go-halo/hook"

	"github.com/fengjx/go-web-quickstart/internal/common"
)

func Init() {
	hook.AddCustomStartHook(common.HookEventBeforeStart, func() {
		initRouter()
	}, common.HookOrderNormal)
}
