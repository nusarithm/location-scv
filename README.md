# Location Service API

API untuk pencarian dan mendapatkan data GeoJSON wilayah Indonesia dengan PostgreSQL + PostGIS.

## 🏗 Arsitektur

```
/location-svc
  /cmd
    main.go              → entrypoint aplikasi
  /internal
    /db                  → koneksi database Postgres + PostGIS
    /models              → struct data (Simple, Location, GeoJSON, dll)
    /repositories        → query ke DB
    /handlers            → Echo handler (search, geojson)
    /routes              → routing endpoint
  /docs                  → auto generated Swagger docs
  docker-compose.yml     → setup database & app
  Dockerfile            → container setup
```

## ⚙️ Teknologi

- **Golang** (Echo framework)
- **PostgreSQL + PostGIS** (DB: `bgn`)
- **Swagger** untuk dokumentasi API
- **Docker & Docker Compose**

## 🚀 Quick Start

### Menggunakan Docker Compose (Recommended)

```bash
# Clone dan masuk ke directory
cd location-svc

# Build dan jalankan semua services
docker-compose up --build

# App akan berjalan di http://localhost:8080
# Swagger docs: http://localhost:8080/swagger/index.html
```

### Manual Setup

1. **Setup Database**
   ```bash
   # Jalankan PostgreSQL + PostGIS
   docker run --name bgn-db -e POSTGRES_DB=bgn -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgis/postgis:15-3.3
   ```

2. **Generate Swagger Docs**
   ```bash
   # Install swag
   go install github.com/swaggo/swag/cmd/swag@latest
   
   # Generate docs
   swag init -g cmd/main.go
   ```

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

4. **Run Application**
   ```bash
   go run cmd/main.go
   ```

## 📖 API Documentation

### 🔎 Search Endpoints (Tag: `search`)

| Method | Endpoint | Description | Query Params |
|--------|----------|-------------|--------------|
| GET | `/search/propinsi` | List semua propinsi | - |
| GET | `/search/kabupaten` | List kabupaten dalam propinsi | `propinsi_id` |
| GET | `/search/kecamatan` | List kecamatan dalam kabupaten | `kabupaten_id` |
| GET | `/search/kelurahan` | List kelurahan dalam kecamatan | `kecamatan_id` |

**Response Format:**
```json
[
  { "id": 3201, "name": "Kabupaten Bogor" },
  { "id": 3202, "name": "Kabupaten Sukabumi" }
]
```

### 🌍 GeoJSON Endpoints (Tag: `geojson`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/geojson/propinsi/:id` | GeoJSON data propinsi |
| GET | `/geojson/kabupaten/:id` | GeoJSON data kabupaten |
| GET | `/geojson/kecamatan/:id` | GeoJSON data kecamatan |
| GET | `/geojson/kelurahan/:id` | GeoJSON data kelurahan |

**Response Format:**
```json
{
  "type": "Feature",
  "properties": {
    "id": 3201,
    "name": "Kabupaten Bogor",
    "type": "kabupaten"
  },
  "geometry": {
    "type": "Polygon",
    "coordinates": [...]
  }
}
```

## 🔧 Environment Variables

| Variable | Default Value | Description |
|----------|---------------|-------------|
| `DATABASE_URL` | `postgres://postgres:postgres@localhost:5432/bgn?sslmode=disable` | PostgreSQL connection string |

## 📝 Development

### Regenerate Swagger Documentation

```bash
swag init -g cmd/main.go
```

### Database Schema

Service ini mengasumsikan tabel berikut sudah ada di database `bgn`:

- `propinsi` (id, name, geom)
- `kabupaten` (id, name, propinsi_id, geom)
- `kecamatan` (id, name, kabupaten_id, geom)
- `kelurahan` (id, name, kecamatan_id, geom)

Semua tabel memiliki kolom `geom` dengan tipe PostGIS geometry.

### Testing Endpoints

```bash
# Test health check
curl http://localhost:8080/health

# Test search propinsi
curl http://localhost:8080/search/propinsi

# Test search kabupaten
curl "http://localhost:8080/search/kabupaten?propinsi_id=32"

# Test GeoJSON
curl http://localhost:8080/geojson/propinsi/32
```

## 📊 Swagger Documentation

Setelah aplikasi berjalan, akses Swagger UI di:
**http://localhost:8080/swagger/index.html**

## 🐳 Docker Commands

```bash
# Build image
docker build -t location-svc .

# Run with Docker Compose
docker-compose up --build

# Stop services
docker-compose down

# View logs
docker-compose logs -f app
```

## 🔑 Aturan Coding

- Semua query ke DB lewat **repositories**
- Handler hanya untuk `bind request` & `return response`
- Gunakan `ST_AsGeoJSON` untuk geometry → frontend Mapbox
- Pastikan ada `@Tags` pada anotasi Swagger untuk grouping
- Gunakan `snake_case` untuk nama kolom DB, `camelCase` untuk JSON response