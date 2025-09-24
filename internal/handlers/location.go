package handlers

import (
	"location-svc/internal/models"
	"location-svc/internal/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LocationHandler struct {
	repo *repositories.LocationRepository
}

// NewLocationHandler creates new instance of LocationHandler
func NewLocationHandler(repo *repositories.LocationRepository) *LocationHandler {
	return &LocationHandler{repo: repo}
}

// GetPropinsi godoc
// @Summary Get all provinces
// @Description Get list of all provinces in Indonesia
// @Tags search
// @Accept json
// @Produce json
// @Param name query string false "Search by province name"
// @Success 200 {array} models.Propinsi
// @Router /search/propinsi [get]
func (h *LocationHandler) GetPropinsi(c echo.Context) error {
	name := c.QueryParam("name")

	var provinces []models.Propinsi
	var err error

	if name != "" {
		provinces, err = h.repo.SearchPropinsiByName(name)
	} else {
		provinces, err = h.repo.GetPropinsi()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get provinces",
		})
	}

	return c.JSON(http.StatusOK, provinces)
}

// GetKabupaten godoc
// @Summary Get regencies by province ID
// @Description Get list of regencies/cities in a specific province
// @Tags search
// @Accept json
// @Produce json
// @Param propinsi_id query string true "Province ID"
// @Param name query string false "Search by regency name"
// @Success 200 {array} models.Kabupaten
// @Router /search/kabupaten [get]
func (h *LocationHandler) GetKabupaten(c echo.Context) error {
	propinsiIDStr := c.QueryParam("propinsi_id")
	name := c.QueryParam("name")

	if propinsiIDStr == "" && name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "propinsi_id is required when not searching by name",
		})
	}

	var kabupatens []models.Kabupaten
	var err error

	if name != "" {
		var propinsiID *string
		if propinsiIDStr != "" {
			propinsiID = &propinsiIDStr
		}
		kabupatens, err = h.repo.SearchKabupatenByName(name, propinsiID)
	} else {
		kabupatens, err = h.repo.GetKabupaten(propinsiIDStr)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get regencies",
		})
	}

	return c.JSON(http.StatusOK, kabupatens)
}

// GetKecamatan godoc
// @Summary Get districts by province and regency ID
// @Description Get list of districts in a specific regency
// @Tags search
// @Accept json
// @Produce json
// @Param propinsi_id query string true "Province ID"
// @Param kabupaten_id query string true "Regency ID"
// @Param name query string false "Search by district name"
// @Success 200 {array} models.Kecamatan
// @Router /search/kecamatan [get]
func (h *LocationHandler) GetKecamatan(c echo.Context) error {
	propinsiIDStr := c.QueryParam("propinsi_id")
	kabupatenIDStr := c.QueryParam("kabupaten_id")
	name := c.QueryParam("name")

	if name == "" && (propinsiIDStr == "" || kabupatenIDStr == "") {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "propinsi_id and kabupaten_id are required when not searching by name",
		})
	}

	var kecamatans []models.Kecamatan
	var err error

	if name != "" {
		var propinsiID, kabupatenID *string

		if propinsiIDStr != "" {
			propinsiID = &propinsiIDStr
		}

		if kabupatenIDStr != "" {
			kabupatenID = &kabupatenIDStr
		}

		kecamatans, err = h.repo.SearchKecamatanByName(name, propinsiID, kabupatenID)
	} else {
		kecamatans, err = h.repo.GetKecamatan(propinsiIDStr, kabupatenIDStr)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get districts",
		})
	}

	return c.JSON(http.StatusOK, kecamatans)
}

// GetKelurahan godoc
// @Summary Get villages by full hierarchy
// @Description Get list of villages in a specific district with full hierarchy
// @Tags search
// @Accept json
// @Produce json
// @Param propinsi_id query string true "Province ID"
// @Param kabupaten_id query string true "Regency ID"
// @Param kecamatan_id query string true "District ID"
// @Param name query string false "Search by village name"
// @Success 200 {array} models.Kelurahan
// @Router /search/kelurahan [get]
func (h *LocationHandler) GetKelurahan(c echo.Context) error {
	propinsiIDStr := c.QueryParam("propinsi_id")
	kabupatenIDStr := c.QueryParam("kabupaten_id")
	kecamatanIDStr := c.QueryParam("kecamatan_id")
	name := c.QueryParam("name")

	if name == "" && (propinsiIDStr == "" || kabupatenIDStr == "" || kecamatanIDStr == "") {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "propinsi_id, kabupaten_id, and kecamatan_id are required when not searching by name",
		})
	}

	var kelurahans []models.Kelurahan
	var err error

	if name != "" {
		var propinsiID, kabupatenID, kecamatanID *string

		if propinsiIDStr != "" {
			propinsiID = &propinsiIDStr
		}

		if kabupatenIDStr != "" {
			kabupatenID = &kabupatenIDStr
		}

		if kecamatanIDStr != "" {
			kecamatanID = &kecamatanIDStr
		}

		kelurahans, err = h.repo.SearchKelurahanByName(name, propinsiID, kabupatenID, kecamatanID)
	} else {
		kelurahans, err = h.repo.GetKelurahan(propinsiIDStr, kabupatenIDStr, kecamatanIDStr)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get villages",
		})
	}

	return c.JSON(http.StatusOK, kelurahans)
}

// GetPropinsiGeoJSON godoc
// @Summary Get province GeoJSON
// @Description Get GeoJSON data for a specific province
// @Tags geojson
// @Accept json
// @Produce json
// @Param id path string true "Province ID"
// @Success 200 {object} models.GeoJSONFeature
// @Router /geojson/propinsi/{id} [get]
func (h *LocationHandler) GetPropinsiGeoJSON(c echo.Context) error {
	id := c.Param("id")

	geojson, err := h.repo.GetPropinsiGeoJSON(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get province GeoJSON",
		})
	}

	return c.JSON(http.StatusOK, geojson)
}

// GetKabupatenGeoJSON godoc
// @Summary Get regency GeoJSON
// @Description Get GeoJSON data for a specific regency
// @Tags geojson
// @Accept json
// @Produce json
// @Param id path string true "Regency ID"
// @Success 200 {object} models.GeoJSONFeature
// @Router /geojson/kabupaten/{id} [get]
func (h *LocationHandler) GetKabupatenGeoJSON(c echo.Context) error {
	id := c.Param("id")

	geojson, err := h.repo.GetKabupatenGeoJSON(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get regency GeoJSON",
		})
	}

	return c.JSON(http.StatusOK, geojson)
}

// GetKecamatanGeoJSON godoc
// @Summary Get district GeoJSON
// @Description Get GeoJSON data for a specific district
// @Tags geojson
// @Accept json
// @Produce json
// @Param id path string true "District ID"
// @Success 200 {object} models.GeoJSONFeature
// @Router /geojson/kecamatan/{id} [get]
func (h *LocationHandler) GetKecamatanGeoJSON(c echo.Context) error {
	id := c.Param("id")

	geojson, err := h.repo.GetKecamatanGeoJSON(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get district GeoJSON",
		})
	}

	return c.JSON(http.StatusOK, geojson)
}

// GetKelurahanGeoJSON godoc
// @Summary Get village GeoJSON
// @Description Get GeoJSON data for a specific village
// @Tags geojson
// @Accept json
// @Produce json
// @Param id path string true "Village ID"
// @Success 200 {object} models.GeoJSONFeature
// @Router /geojson/kelurahan/{id} [get]
func (h *LocationHandler) GetKelurahanGeoJSON(c echo.Context) error {
	id := c.Param("id")

	geojson, err := h.repo.GetKelurahanGeoJSON(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get village GeoJSON",
		})
	}

	return c.JSON(http.StatusOK, geojson)
}
