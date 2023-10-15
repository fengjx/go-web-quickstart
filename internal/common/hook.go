package common

import "math"

const (
	HookOrderHighest = 1
	HookOrderNormal  = 100
	HookOrderLowest  = math.MaxInt
)

const (
	HookEventBeforeStart = "before-start"
	HookEventShutdown    = "shutdown"
)
