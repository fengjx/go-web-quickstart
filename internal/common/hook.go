package common

import "math"

const (
	HookOrderHighest = 1
	HookOrderNormal  = 100
	HookOrderLowest  = math.MaxInt
)

const (
	HookEventEndpointInitFinish = "endpoint-init-finish"
	HookEventBeforeStart        = "before-start"
	HookEventAfterStart         = "after-start"
	HookEventBeforeShutdown     = "before-shutdown"
	HookEventAfterShutdown      = "after-shutdown"
)
