package hook

import (
	"math"
	"sort"
	"sync"

	"github.com/samber/lo"
)

const (
	// OrderHighest 越小优先级越高
	OrderHighest = 1           // 配置预加载
	Order10      = 10          // 配置预加载
	Order100     = 100         // service 依赖顺序控制
	Order1000    = 1000        // service 依赖顺序控制
	OrderLowest  = math.MaxInt // 最迟加载，如kafka消费者注册，本地定时任务，延迟队列等
)

// order 越小优先级越高
type hookFun struct {
	handler func()
	order   int
}

var startHooks []hookFun
var startHooksLock sync.Mutex

var stopHooks []hookFun
var stopHooksLock sync.Mutex

func AddStartHook(handler func(), order int) {
	startHooksLock.Lock()
	defer startHooksLock.Unlock()
	startHooks = append(startHooks, hookFun{
		handler: handler,
		order:   order,
	})
}

func AddStopHook(handler func(), order int) {
	stopHooksLock.Lock()
	defer stopHooksLock.Unlock()
	stopHooks = append(stopHooks, hookFun{
		handler: handler,
		order:   order,
	})
}

func OnStart() {
	doHooks(startHooks)
}

func OnStop() {
	doHooks(stopHooks)
}

func doHooks(hookFns []hookFun) {
	hookGroup := make(map[int][]hookFun)
	for _, hook := range hookFns {
		fnList := hookGroup[hook.order]
		hookGroup[hook.order] = append(fnList, hook)
	}
	keys := lo.Keys[int, []hookFun](hookGroup)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, order := range keys {
		hooks := hookGroup[order]
		group := &sync.WaitGroup{}
		group.Add(len(hooks))
		execHooks(hooks, group)
		group.Wait()
	}
}

func execHooks(hooks []hookFun, wg *sync.WaitGroup) {
	for _, fn := range hooks {
		f := fn
		go func() {
			defer wg.Done()
			f.handler()
		}()
	}
}
