-- Create database schema for Indonesian administrative regions
-- This file contains the table structure for Location Service

-- Enable PostGIS extension
CREATE EXTENSION IF NOT EXISTS postgis;

-- Create propinsi table
CREATE TABLE IF NOT EXISTS propinsi (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    geom GEOMETRY(MULTIPOLYGON, 4326)
);

-- Create kabupaten table
CREATE TABLE IF NOT EXISTS kabupaten (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    propinsi_id INTEGER NOT NULL,
    geom GEOMETRY(MULTIPOLYGON, 4326),
    FOREIGN KEY (propinsi_id) REFERENCES propinsi(id)
);

-- Create kecamatan table
CREATE TABLE IF NOT EXISTS kecamatan (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    kabupaten_id INTEGER NOT NULL,
    geom GEOMETRY(MULTIPOLYGON, 4326),
    FOREIGN KEY (kabupaten_id) REFERENCES kabupaten(id)
);

-- Create kelurahan table
CREATE TABLE IF NOT EXISTS kelurahan (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    kecamatan_id INTEGER NOT NULL,
    geom GEOMETRY(MULTIPOLYGON, 4326),
    FOREIGN KEY (kecamatan_id) REFERENCES kecamatan(id)
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_kabupaten_propinsi_id ON kabupaten(propinsi_id);
CREATE INDEX IF NOT EXISTS idx_kecamatan_kabupaten_id ON kecamatan(kabupaten_id);
CREATE INDEX IF NOT EXISTS idx_kelurahan_kecamatan_id ON kelurahan(kecamatan_id);

-- Create spatial indexes
CREATE INDEX IF NOT EXISTS idx_propinsi_geom ON propinsi USING GIST(geom);
CREATE INDEX IF NOT EXISTS idx_kabupaten_geom ON kabupaten USING GIST(geom);
CREATE INDEX IF NOT EXISTS idx_kecamatan_geom ON kecamatan USING GIST(geom);
CREATE INDEX IF NOT EXISTS idx_kelurahan_geom ON kelurahan USING GIST(geom);