package main

import (
	"location-svc/internal/db"
	"location-svc/internal/routes"
	"log"
	"os"

	_ "location-svc/docs" // Import untuk swagger docs

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Location Service API
// @version 1.0
// @description API untuk pencarian dan mendapatkan data GeoJSON wilayah Indonesia
// @host location-svc.nusarithm.id
// @BasePath /
func main() {
	// Inisialisasi Echo
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded:", err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"*"},
	}))

	// Inisialisasi koneksi database
	database, err := db.Init()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Setup routes
	routes.Setup(e, database)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Println("Server starting on :" + PORT)
	e.Logger.Fatal(e.Start(":" + PORT))
}
