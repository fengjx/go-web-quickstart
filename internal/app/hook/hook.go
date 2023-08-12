package hook

type hookFun func()

var startHooks []hookFun
var stopHooks []hookFun

func AddStartHook(fun hookFun) {
	startHooks = append(startHooks, fun)
}

func AddStopHook(fun hookFun) {
	stopHooks = append(stopHooks, fun)
}

func OnStart() {
	for _, hook := range startHooks {
		hook()
	}
}

func OnStop() {
	for _, hook := range stopHooks {
		hook()
	}
}
