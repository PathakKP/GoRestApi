package routes

import (
	"GoRestApi/controller"
	"GoRestApi/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
)

// MovieRoutes initializes all the routes and returns a router
func MovieRoutes() *mux.Router {
	var router = mux.NewRouter().StrictSlash(true)

	// Home Route
	router.HandleFunc("/api/", func(rw http.ResponseWriter, r *http.Request) {
		logRequestGID("Home Route")
		message := models.Message{
			Message: "Movie API!!!!",
		}
		json.NewEncoder(rw).Encode(message)
	})

	// Other Routes with goroutine tracking
	router.HandleFunc("/api/movies", wrapWithGID(controller.AddMovie)).Methods(http.MethodPost)
	router.HandleFunc("/api/movies", wrapWithGID(controller.GetAllMovies)).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}", wrapWithGID(controller.GetMovieByID)).Methods(http.MethodGet)
	// router.HandleFunc("/api/movies/{id}", wrapWithGID(controller.DeleteMovieById)).Methods(http.MethodDelete)
	router.HandleFunc("/api/movies", wrapWithGID(controller.UpdateMovie)).Methods(http.MethodPut)

	return router
}

// Helper function to log the current GID with the route name
func logRequestGID(route string) {
	gid := getGID()
	log.Printf("[GID: %d] Handling %s\n", gid, route)
}

// Middleware-like wrapper to log GID for a route handler
func wrapWithGID(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logRequestGID(r.URL.Path) // Log GID for the current request
		handlerFunc(w, r)         // Call the actual handler
	}
}

// getGID retrieves the goroutine ID
func getGID() uint64 {
	b := make([]byte, 64)                            // Create a byte slice with 64 bytes.
	n := runtime.Stack(b, false)                     // Get the current goroutine's stack trace.
	var gid uint64                                   // Variable to store the goroutine ID.
	fmt.Sscanf(string(b[:n]), "goroutine %d ", &gid) // Extract the goroutine ID from the stack trace.
	return gid                                       // Return the goroutine ID.
}
