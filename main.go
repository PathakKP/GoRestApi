package main

import (
	"GoRestApi/db"
	"GoRestApi/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Connect to the database
	db.Connect()
	defer db.Close()

	router := routes.MovieRoutes()

	http.Handle("/api", router)

	log.Println("Listening on port 8081...")
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	// Start server in a goroutine to handle shutdown signals
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for an interrupt signal for graceful shutdown
	waitForShutdown(server)

}

// Helper function to gracefully shut down the server
func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully.")
}
