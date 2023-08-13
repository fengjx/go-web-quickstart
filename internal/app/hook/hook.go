package hook

import (
	"sort"
)

// order 越小优先级越高
type hookFun struct {
	handler func()
	order   int
}

var startHooks []hookFun
var stopHooks []hookFun

func AddStartHook(handler func(), order int) {
	startHooks = append(startHooks, hookFun{
		handler: handler,
		order:   order,
	})
	sort.Slice(startHooks, func(i, j int) bool {
		return startHooks[i].order < startHooks[j].order
	})
}

func AddStopHook(handler func(), order int) {
	stopHooks = append(stopHooks, hookFun{
		handler: handler,
		order:   order,
	})
	sort.Slice(stopHooks, func(i, j int) bool {
		return stopHooks[i].order < stopHooks[j].order
	})
}

func OnStart() {
	for _, hook := range startHooks {
		hook.handler()
	}
}

func OnStop() {
	for _, hook := range stopHooks {
		hook.handler()
	}
}
