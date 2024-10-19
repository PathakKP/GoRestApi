package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var DB *sql.DB

// Connect initializes the PostgreSQL connection
func Connect() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get DB credentials from environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Construct the DSN (Database Source Name) for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a new database connection
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to open database: %v", err)
	}

	// Verify the connection is valid
	if err := DB.Ping(); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	fmt.Println("Connected to Postgres!")

	// Run database migrations
	runMigrations(DB)
}

// runMigrations applies any pending migrations
func runMigrations(db *sql.DB) {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

// Close closes the database connection
func Close() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection closed.")
	}
}
