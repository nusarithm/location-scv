package routes

import (
	"database/sql"
	"location-svc/internal/handlers"
	"location-svc/internal/repositories"

	"github.com/labstack/echo/v4"
)

// Setup menginisialisasi semua routes untuk aplikasi
func Setup(e *echo.Echo, db *sql.DB) {
	// Initialize repository
	locationRepo := repositories.NewLocationRepository(db)

	// Initialize handler
	locationHandler := handlers.NewLocationHandler(locationRepo)

	// Search endpoints (Tag: search)
	searchGroup := e.Group("/search")
	searchGroup.GET("/propinsi", locationHandler.GetPropinsi)
	searchGroup.GET("/kabupaten", locationHandler.GetKabupaten)
	searchGroup.GET("/kecamatan", locationHandler.GetKecamatan)
	searchGroup.GET("/kelurahan", locationHandler.GetKelurahan)

	// GeoJSON endpoints (Tag: geojson)
	geojsonGroup := e.Group("/geojson")
	geojsonGroup.GET("/propinsi/:id", locationHandler.GetPropinsiGeoJSON)
	geojsonGroup.GET("/kabupaten/:id", locationHandler.GetKabupatenGeoJSON)
	geojsonGroup.GET("/kecamatan/:id", locationHandler.GetKecamatanGeoJSON)
	geojsonGroup.GET("/kelurahan/:id", locationHandler.GetKelurahanGeoJSON)

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":  "OK",
			"service": "Location Service",
		})
	})
}
