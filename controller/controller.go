package controller

import (
	"GoRestApi/models"
	"GoRestApi/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllMovies handles the request to retrieve all movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := repository.GetAllMovies()
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

// GetMovieByID handles the request to retrieve a specific movie by ID
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	movie, err := repository.GetMovieByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

// AddMovie handles the request to add a new movie
func AddMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := repository.AddMovie(movie); err != nil {
		log.Printf("Failed to add movie: %v", err)
		http.Error(w, "Failed to add movie", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models.Message{Message: "Movie added"})
}

// DeleteMovieByID handles the request to delete a movie by ID
func DeleteMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := repository.DeleteMovieByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(models.Message{Message: "Movie deleted"})
}

// UpdateMovie handles the request to update an existing movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := repository.UpdateMovie(movie); err != nil {
		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models.Message{Message: "Movie updated"})
}
