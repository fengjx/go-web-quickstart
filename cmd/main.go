package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengjx/go-web-quickstart/internal/app"
	"github.com/fengjx/go-web-quickstart/internal/app/appconfig"
	"github.com/fengjx/go-web-quickstart/internal/common/env"
)

func main() {
	log.Println("app start env:", env.GetEnv())
	appconfig.Init()
	ctx, cancel := context.WithCancel(context.Background())
	app.Start(ctx)
	// Wait for signal to initiate server shutdown.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	cancel()
	app.Stop(ctx)
}
