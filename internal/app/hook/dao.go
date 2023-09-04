package hook

import "sync"

var daoInitHooks []hookFun
var daoInitHooksLock sync.Mutex

func AddDaoInitHook(handler func(), order int) {
	daoInitHooksLock.Lock()
	defer daoInitHooksLock.Unlock()
	daoInitHooks = append(daoInitHooks, hookFun{
		handler: handler,
		order:   order,
	})
}

func OnDaoInit() {
	doHooks(daoInitHooks)
}
