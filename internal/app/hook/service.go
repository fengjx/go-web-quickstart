package hook

import "sync"

var serviceInitHooks []hookFun
var serviceInitHooksLock sync.Mutex

func AddServiceInitHooksHook(handler func(), order int) {
	serviceInitHooksLock.Lock()
	defer serviceInitHooksLock.Unlock()
	serviceInitHooks = append(serviceInitHooks, hookFun{
		handler: handler,
		order:   order,
	})
}

func OnServiceInit() {
	doHooks(serviceInitHooks)
}
