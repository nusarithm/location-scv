package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"location-svc/internal/models"
)

type LocationRepository struct {
	db *sql.DB
}

// NewLocationRepository creates new instance of LocationRepository
func NewLocationRepository(db *sql.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

// GetPropinsi mendapatkan semua propinsi
func (r *LocationRepository) GetPropinsi() ([]models.Propinsi, error) {
	query := `SELECT kd_propinsi, nm_propinsi FROM propinsi ORDER BY nm_propinsi`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var provinces []models.Propinsi
	for rows.Next() {
		var p models.Propinsi
		if err := rows.Scan(&p.KdPropinsi, &p.NmPropinsi); err != nil {
			return nil, err
		}
		provinces = append(provinces, p)
	}

	return provinces, nil
}

// SearchPropinsiByName mencari propinsi berdasarkan nama
func (r *LocationRepository) SearchPropinsiByName(name string) ([]models.Propinsi, error) {
	query := `SELECT kd_propinsi, nm_propinsi FROM propinsi WHERE nm_propinsi ILIKE '%' || $1 || '%' ORDER BY nm_propinsi`

	rows, err := r.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var provinces []models.Propinsi
	for rows.Next() {
		var p models.Propinsi
		if err := rows.Scan(&p.KdPropinsi, &p.NmPropinsi); err != nil {
			return nil, err
		}
		provinces = append(provinces, p)
	}

	return provinces, nil
}

// GetKabupaten mendapatkan kabupaten berdasarkan propinsi_id
func (r *LocationRepository) GetKabupaten(propinsiID string) ([]models.Kabupaten, error) {
	query := `
		SELECT p.kd_propinsi, p.nm_propinsi, k.kd_kabupaten, k.nm_kabupaten 
		FROM kabupaten k 
		JOIN propinsi p ON k.kd_propinsi = p.kd_propinsi 
		WHERE k.kd_propinsi = $1 
		ORDER BY k.nm_kabupaten`

	rows, err := r.db.Query(query, propinsiID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kabupatens []models.Kabupaten
	for rows.Next() {
		var k models.Kabupaten
		if err := rows.Scan(&k.KdPropinsi, &k.NmPropinsi, &k.KdKabupaten, &k.NmKabupaten); err != nil {
			return nil, err
		}
		kabupatens = append(kabupatens, k)
	}

	return kabupatens, nil
}

// SearchKabupatenByName mencari kabupaten berdasarkan nama dengan filter propinsi
func (r *LocationRepository) SearchKabupatenByName(name string, propinsiID *string) ([]models.Kabupaten, error) {
	var query string
	var args []interface{}

	if propinsiID != nil {
		query = `
			SELECT p.kd_propinsi, p.nm_propinsi, k.kd_kabupaten, k.nm_kabupaten 
			FROM kabupaten k 
			JOIN propinsi p ON k.kd_propinsi = p.kd_propinsi 
			WHERE k.nm_kabupaten ILIKE '%' || $1 || '%' AND k.kd_propinsi = $2
			ORDER BY k.nm_kabupaten`
		args = []interface{}{name, *propinsiID}
	} else {
		query = `
			SELECT p.kd_propinsi, p.nm_propinsi, k.kd_kabupaten, k.nm_kabupaten 
			FROM kabupaten k 
			JOIN propinsi p ON k.kd_propinsi = p.kd_propinsi 
			WHERE k.nm_kabupaten ILIKE '%' || $1 || '%'
			ORDER BY k.nm_kabupaten`
		args = []interface{}{name}
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kabupatens []models.Kabupaten
	for rows.Next() {
		var k models.Kabupaten
		if err := rows.Scan(&k.KdPropinsi, &k.NmPropinsi, &k.KdKabupaten, &k.NmKabupaten); err != nil {
			return nil, err
		}
		kabupatens = append(kabupatens, k)
	}

	return kabupatens, nil
} // GetKecamatan mendapatkan kecamatan berdasarkan propinsi_id dan kabupaten_id
func (r *LocationRepository) GetKecamatan(propinsiID, kabupatenID string) ([]models.Kecamatan, error) {
	query := `
		SELECT p.kd_propinsi, p.nm_propinsi, k.kd_kabupaten, k.nm_kabupaten, kec.kd_kecamatan, kec.nm_kecamatan
		FROM kecamatan kec
		JOIN kabupaten k ON kec.kd_kabupaten = k.kd_kabupaten
		JOIN propinsi p ON k.kd_propinsi = p.kd_propinsi
		WHERE k.kd_propinsi = $1 AND k.kd_kabupaten = $2
		ORDER BY kec.nm_kecamatan`

	rows, err := r.db.Query(query, propinsiID, kabupatenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kecamatans []models.Kecamatan
	for rows.Next() {
		var kec models.Kecamatan
		if err := rows.Scan(&kec.KdPropinsi, &kec.NmPropinsi, &kec.KdKabupaten, &kec.NmKabupaten, &kec.KdKecamatan, &kec.NmKecamatan); err != nil {
			return nil, err
		}
		kecamatans = append(kecamatans, kec)
	}

	return kecamatans, nil
}

// SearchKecamatanByName mencari kecamatan berdasarkan nama dengan filter hierarki
func (r *LocationRepository) SearchKecamatanByName(name string, propinsiID, kabupatenID *string) ([]models.Kecamatan, error) {
	var query string
	var args []interface{}

	baseQuery := `
		SELECT p.kd_propinsi, p.nm_propinsi, k.kd_kabupaten, k.nm_kabupaten, kec.kd_kecamatan, kec.nm_kecamatan
		FROM kecamatan kec
		JOIN kabupaten k ON kec.kd_kabupaten = k.kd_kabupaten
		JOIN propinsi p ON k.kd_propinsi = p.kd_propinsi
		WHERE kec.nm_kecamatan ILIKE '%' || $1 || '%'`

	args = append(args, name)

	if propinsiID != nil {
		baseQuery += " AND k.kd_propinsi = $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, *propinsiID)
	}

	if kabupatenID != nil {
		baseQuery += " AND k.kd_kabupaten = $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, *kabupatenID)
	}

	query = baseQuery + " ORDER BY kec.nm_kecamatan"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kecamatans []models.Kecamatan
	for rows.Next() {
		var kec models.Kecamatan
		if err := rows.Scan(&kec.KdPropinsi, &kec.NmPropinsi, &kec.KdKabupaten, &kec.NmKabupaten, &kec.KdKecamatan, &kec.NmKecamatan); err != nil {
			return nil, err
		}
		kecamatans = append(kecamatans, kec)
	}

	return kecamatans, nil
} // GetKelurahan mendapatkan kelurahan berdasarkan hierarki lengkap
func (r *LocationRepository) GetKelurahan(propinsiID, kabupatenID, kecamatanID string) ([]models.Kelurahan, error) {
	query := `
		SELECT p.kd_propinsi, p.nm_propinsi, k.kd_kabupaten, k.nm_kabupaten, 
		       kec.kd_kecamatan, kec.nm_kecamatan, kel.kd_kelurahan, kel.nm_kelurahan
		FROM kelurahan kel
		JOIN kecamatan kec ON kel.kd_kecamatan = kec.kd_kecamatan
		JOIN kabupaten k ON kec.kd_kabupaten = k.kd_kabupaten
		JOIN propinsi p ON k.kd_propinsi = p.kd_propinsi
		WHERE p.kd_propinsi = $1 AND k.kd_kabupaten = $2 AND kec.kd_kecamatan = $3
		ORDER BY kel.nm_kelurahan`

	rows, err := r.db.Query(query, propinsiID, kabupatenID, kecamatanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kelurahans []models.Kelurahan
	for rows.Next() {
		var kel models.Kelurahan
		if err := rows.Scan(&kel.KdPropinsi, &kel.NmPropinsi, &kel.KdKabupaten, &kel.NmKabupaten,
			&kel.KdKecamatan, &kel.NmKecamatan, &kel.KdKelurahan, &kel.NmKelurahan); err != nil {
			return nil, err
		}
		kelurahans = append(kelurahans, kel)
	}

	return kelurahans, nil
}

// SearchKelurahanByName mencari kelurahan berdasarkan nama dengan filter hierarki
func (r *LocationRepository) SearchKelurahanByName(name string, propinsiID, kabupatenID, kecamatanID *string) ([]models.Kelurahan, error) {
	var query string
	var args []interface{}

	baseQuery := `
		SELECT p.kd_propinsi, p.nm_propinsi, k.kd_kabupaten, k.nm_kabupaten, 
		       kec.kd_kecamatan, kec.nm_kecamatan, kel.kd_kelurahan, kel.nm_kelurahan
		FROM kelurahan kel
		JOIN kecamatan kec ON kel.kd_kecamatan = kec.kd_kecamatan
		JOIN kabupaten k ON kec.kd_kabupaten = k.kd_kabupaten
		JOIN propinsi p ON k.kd_propinsi = p.kd_propinsi
		WHERE kel.nm_kelurahan ILIKE '%' || $1 || '%'`

	args = append(args, name)

	if propinsiID != nil {
		baseQuery += " AND p.kd_propinsi = $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, *propinsiID)
	}

	if kabupatenID != nil {
		baseQuery += " AND k.kd_kabupaten = $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, *kabupatenID)
	}

	if kecamatanID != nil {
		baseQuery += " AND kec.kd_kecamatan = $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, *kecamatanID)
	}

	query = baseQuery + " ORDER BY kel.nm_kelurahan"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kelurahans []models.Kelurahan
	for rows.Next() {
		var kel models.Kelurahan
		if err := rows.Scan(&kel.KdPropinsi, &kel.NmPropinsi, &kel.KdKabupaten, &kel.NmKabupaten,
			&kel.KdKecamatan, &kel.NmKecamatan, &kel.KdKelurahan, &kel.NmKelurahan); err != nil {
			return nil, err
		}
		kelurahans = append(kelurahans, kel)
	}

	return kelurahans, nil
}

// GetPropinsiGeoJSON mendapatkan GeoJSON propinsi berdasarkan ID
func (r *LocationRepository) GetPropinsiGeoJSON(id string) (*models.GeoJSONFeature, error) {
	query := `
		SELECT kd_propinsi, nm_propinsi, ST_AsGeoJSON(geom) as geometry 
		FROM propinsi 
		WHERE kd_propinsi = $1
	`

	var feature models.GeoJSONFeature
	var geometryJSON string

	err := r.db.QueryRow(query, id).Scan(&feature.Properties.ID, &feature.Properties.Name, &geometryJSON)
	if err != nil {
		return nil, err
	}

	feature.Type = "Feature"
	feature.Properties.Type = "propinsi"

	if err := json.Unmarshal([]byte(geometryJSON), &feature.Geometry); err != nil {
		return nil, err
	}

	return &feature, nil
}

// GetKabupatenGeoJSON mendapatkan GeoJSON kabupaten berdasarkan ID
func (r *LocationRepository) GetKabupatenGeoJSON(id string) (*models.GeoJSONFeature, error) {
	query := `
		SELECT kd_kabupaten, nm_kabupaten, ST_AsGeoJSON(geom) as geometry 
		FROM kabupaten 
		WHERE kd_kabupaten = $1
	`

	var feature models.GeoJSONFeature
	var geometryJSON string

	err := r.db.QueryRow(query, id).Scan(&feature.Properties.ID, &feature.Properties.Name, &geometryJSON)
	if err != nil {
		return nil, err
	}

	feature.Type = "Feature"
	feature.Properties.Type = "kabupaten"

	if err := json.Unmarshal([]byte(geometryJSON), &feature.Geometry); err != nil {
		return nil, err
	}

	return &feature, nil
}

// GetKecamatanGeoJSON mendapatkan GeoJSON kecamatan berdasarkan ID
func (r *LocationRepository) GetKecamatanGeoJSON(id string) (*models.GeoJSONFeature, error) {
	query := `
		SELECT kd_kecamatan, nm_kecamatan, ST_AsGeoJSON(geom) as geometry 
		FROM kecamatan 
		WHERE kd_kecamatan = $1
	`

	var feature models.GeoJSONFeature
	var geometryJSON string

	err := r.db.QueryRow(query, id).Scan(&feature.Properties.ID, &feature.Properties.Name, &geometryJSON)
	if err != nil {
		return nil, err
	}

	feature.Type = "Feature"
	feature.Properties.Type = "kecamatan"

	if err := json.Unmarshal([]byte(geometryJSON), &feature.Geometry); err != nil {
		return nil, err
	}

	return &feature, nil
}

// GetKelurahanGeoJSON mendapatkan GeoJSON kelurahan berdasarkan ID
func (r *LocationRepository) GetKelurahanGeoJSON(id string) (*models.GeoJSONFeature, error) {
	query := `
		SELECT kd_kelurahan, nm_kelurahan, ST_AsGeoJSON(geom) as geometry 
		FROM kelurahan 
		WHERE kd_kelurahan = $1
	`

	var feature models.GeoJSONFeature
	var geometryJSON string

	err := r.db.QueryRow(query, id).Scan(&feature.Properties.ID, &feature.Properties.Name, &geometryJSON)
	if err != nil {
		return nil, err
	}

	feature.Type = "Feature"
	feature.Properties.Type = "kelurahan"

	if err := json.Unmarshal([]byte(geometryJSON), &feature.Geometry); err != nil {
		return nil, err
	}

	return &feature, nil
}
