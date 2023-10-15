package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengjx/go-web-quickstart/internal/app"
)

func main() {
	ctx := context.Background()
	app.Start(ctx)
	// Wait for signal to initiate server shutdown.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx.Done()
	app.Stop(ctx)
}
