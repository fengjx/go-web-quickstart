package service

import "context"

type Service interface {
	Init()
	Start(ctx context.Context)
	Shutdown(ctx context.Context)
}
