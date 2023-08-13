package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fengjx/go-web-quickstart/internal/app"
)

func main() {
	app.Start()
	// Wait for signal to initiate server shutdown.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	app.Stop()
}
