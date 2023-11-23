package service

import "context"

type Service interface {
	Start(ctx context.Context)
	Shutdown(ctx context.Context)
}
