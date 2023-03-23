package main

import (
	"context"
	"github.com/fengjx/go-web-quickstart/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	server := app.NewServer()
	server.Start(ctx)
	// Wait for signal to initiate server shutdown.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	server.Shutdown()
	cancel()
}
