package main

import (
	"fmt"
	"go-fiber/middleware"
	"go-fiber/router"
	util "go-fiber/util"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {

	fmt.Println("hello")

	app := fiber.New()
	middleware.InitFiberBuiltinMiddleware(app)

	router.InitApiRoutes(app)
	router.InitSwaggerRoutes(app)
	router.InitNotFoundRoutes(app)

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		startServer(app)
	} else {
		startServerWithGracefulShutdown(app)
	}
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func startServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL, _ := util.ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func startServer(a *fiber.App) {
	// Build Fiber connection URL.
	fiberConnURL, _ := util.ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
