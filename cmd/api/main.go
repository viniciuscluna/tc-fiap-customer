package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/viniciuscluna/tc-fiap-customer/docs"

	"github.com/viniciuscluna/tc-fiap-customer/internal/app"
)

// @title           Tc-Fiap-50
// @version         1.0
// @description     Api to manage a fast food restaurant
// @host            localhost:8080
// @BasePath        /
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capture system signals for graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	// Initialize the application using Uber FX
	app := app.InitializeApp()

	// Start the Uber FX lifecycle
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Error while starting app: %v", err)
	}

	// Wait until the context is canceled
	<-ctx.Done()

	// Stop the Uber FX lifecycle
	if err := app.Stop(ctx); err != nil {
		log.Fatalf("Error while stopping app: %v", err)
	}
}
