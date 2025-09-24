package models

// Propinsi represents provinsi data
type Propinsi struct {
	KdPropinsi string `json:"kd_propinsi"`
	NmPropinsi string `json:"nm_propinsi"`
}

// Kabupaten represents kabupaten data with provinsi info
type Kabupaten struct {
	KdPropinsi  string `json:"kd_propinsi"`
	NmPropinsi  string `json:"nm_propinsi"`
	KdKabupaten string `json:"kd_kabupaten"`
	NmKabupaten string `json:"nm_kabupaten"`
}

// Kecamatan represents kecamatan data with full hierarchy
type Kecamatan struct {
	KdPropinsi  string `json:"kd_propinsi"`
	NmPropinsi  string `json:"nm_propinsi"`
	KdKabupaten string `json:"kd_kabupaten"`
	NmKabupaten string `json:"nm_kabupaten"`
	KdKecamatan string `json:"kd_kecamatan"`
	NmKecamatan string `json:"nm_kecamatan"`
}

// Kelurahan represents kelurahan data with full hierarchy
type Kelurahan struct {
	KdPropinsi  string `json:"kd_propinsi"`
	NmPropinsi  string `json:"nm_propinsi"`
	KdKabupaten string `json:"kd_kabupaten"`
	NmKabupaten string `json:"nm_kabupaten"`
	KdKecamatan string `json:"kd_kecamatan"`
	NmKecamatan string `json:"nm_kecamatan"`
	KdKelurahan string `json:"kd_kelurahan"`
	NmKelurahan string `json:"nm_kelurahan"`
}

// GeoJSONFeature represents GeoJSON Feature format
type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Properties GeoJSONProperties      `json:"properties"`
	Geometry   map[string]interface{} `json:"geometry"`
}

// GeoJSONProperties represents properties dalam GeoJSON Feature
type GeoJSONProperties struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
