package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lauslim12/expert-systems/internal/application"
)

// Constant for the web frontend location.
const pathToWebDirectory = "./web/build"

// Get port from environment variable. If it does not exist, use '8080'.
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}

	return fmt.Sprintf(":%s", port)
}

// Get application mode from environment variable 'GO_ENV'. If it does not exist, use 'development'.
func getMode() string {
	mode := os.Getenv("GO_ENV")
	if mode != "production" {
		return "development"
	}

	return "production"
}

// Starting point, initialize server.
func main() {
	// HTTP server initialization.
	server := &http.Server{Addr: getPort(), Handler: application.Configure(pathToWebDirectory, getMode())}

	// Prepare context for graceful shutdown.
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt or quit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds.
		shutdownCtx, shutdownCtxCancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer shutdownCtxCancel()
		log.Println("Server starting to shutdown in 30 seconds...")

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timeout, forcing exit.")
			}
		}()

		// Trigger graceful shutdown here.
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		serverStopCtx()
	}()

	// Run our server and print out starting message.
	log.Printf("Server has started on port %s with environment %s!", getPort(), getMode())
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped.
	<-serverCtx.Done()
	log.Println("Server shut down successfully!")
}
