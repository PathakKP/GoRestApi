package repository

import (
	"GoRestApi/db"
	"GoRestApi/models"
	"context"
	"database/sql"
	"errors"
	"log"
)

// GetAllMovies retrieves all movies from the database
func GetAllMovies() ([]models.Movie, error) {
	rows, err := db.DB.QueryContext(context.Background(), "SELECT id, title, category, year, imdb_rating FROM movies")
	if err != nil {
		log.Println("Error querying movies:", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Category, &movie.Year, &movie.ImdbRating); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// GetMovieByID retrieves a specific movie by ID
func GetMovieByID(id string) (*models.Movie, error) {
	var movie models.Movie
	err := db.DB.QueryRowContext(context.Background(),
		"SELECT id, title, category, year, imdb_rating FROM movies WHERE id = $1", id).
		Scan(&movie.Id, &movie.Title, &movie.Category, &movie.Year, &movie.ImdbRating)

	if err == sql.ErrNoRows {
		return nil, errors.New("movie not found")
	} else if err != nil {
		return nil, err
	}
	return &movie, nil
}

// AddMovie inserts a new movie into the database
func AddMovie(movie models.Movie) error {
	_, err := db.DB.ExecContext(context.Background(),
		"INSERT INTO movies (title, category, year, imdb_rating) VALUES ($1, $2, $3, $4)",
		movie.Title, movie.Category, movie.Year, movie.ImdbRating)
	return err
}

// DeleteMovieByID deletes a movie by ID
func DeleteMovieByID(id string) error {
	result, err := db.DB.ExecContext(context.Background(), "DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no movie found with the given ID")
	}
	return nil
}

// UpdateMovie updates an existing movie
func UpdateMovie(movie models.Movie) error {
	_, err := db.DB.ExecContext(context.Background(),
		"UPDATE movies SET title = $1, category = $2, year = $3, imdb_rating = $4 WHERE id = $5",
		movie.Title, movie.Category, movie.Year, movie.ImdbRating, movie.Id)
	return err
}
