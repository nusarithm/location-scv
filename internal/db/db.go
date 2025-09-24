package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Init menginisialisasi koneksi database PostgreSQL + PostGIS
func Init() (*sql.DB, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded:", err)
	}

	// Mendapatkan database URL dari environment variable atau menggunakan default
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/bgn?sslmode=disable"
	}

	// Membuka koneksi ke database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Test koneksi
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
