-- Sample data for testing Location Service
-- This file contains basic sample data for Indonesian administrative regions

-- Insert sample propinsi data
INSERT INTO propinsi (id, name, geom) VALUES 
(32, 'Jawa Barat', ST_GeomFromText('MULTIPOLYGON(((106.0 -6.0, 109.0 -6.0, 109.0 -8.0, 106.0 -8.0, 106.0 -6.0)))', 4326)),
(31, 'DKI Jakarta', ST_GeomFromText('MULTIPOLYGON(((106.7 -6.1, 106.9 -6.1, 106.9 -6.3, 106.7 -6.3, 106.7 -6.1)))', 4326)),
(33, 'Jawa Tengah', ST_GeomFromText('MULTIPOLYGON(((109.0 -6.0, 112.0 -6.0, 112.0 -8.0, 109.0 -8.0, 109.0 -6.0)))', 4326))
ON CONFLICT (id) DO NOTHING;

-- Insert sample kabupaten data
INSERT INTO kabupaten (id, name, propinsi_id, geom) VALUES 
(3201, 'Kabupaten Bogor', 32, ST_GeomFromText('MULTIPOLYGON(((106.5 -6.3, 107.0 -6.3, 107.0 -6.8, 106.5 -6.8, 106.5 -6.3)))', 4326)),
(3202, 'Kabupaten Sukabumi', 32, ST_GeomFromText('MULTIPOLYGON(((106.8 -6.8, 107.3 -6.8, 107.3 -7.3, 106.8 -7.3, 106.8 -6.8)))', 4326)),
(3171, 'Kota Jakarta Selatan', 31, ST_GeomFromText('MULTIPOLYGON(((106.7 -6.2, 106.85 -6.2, 106.85 -6.35, 106.7 -6.35, 106.7 -6.2)))', 4326)),
(3172, 'Kota Jakarta Timur', 31, ST_GeomFromText('MULTIPOLYGON(((106.85 -6.1, 106.95 -6.1, 106.95 -6.25, 106.85 -6.25, 106.85 -6.1)))', 4326))
ON CONFLICT (id) DO NOTHING;

-- Insert sample kecamatan data
INSERT INTO kecamatan (id, name, kabupaten_id, geom) VALUES 
(3201010, 'Nanggung', 3201, ST_GeomFromText('MULTIPOLYGON(((106.5 -6.3, 106.6 -6.3, 106.6 -6.4, 106.5 -6.4, 106.5 -6.3)))', 4326)),
(3201020, 'Leuwiliang', 3201, ST_GeomFromText('MULTIPOLYGON(((106.6 -6.4, 106.7 -6.4, 106.7 -6.5, 106.6 -6.5, 106.6 -6.4)))', 4326)),
(3202010, 'Pelabuhan Ratu', 3202, ST_GeomFromText('MULTIPOLYGON(((106.8 -6.8, 106.9 -6.8, 106.9 -6.9, 106.8 -6.9, 106.8 -6.8)))', 4326)),
(3171010, 'Jagakarsa', 3171, ST_GeomFromText('MULTIPOLYGON(((106.7 -6.2, 106.75 -6.2, 106.75 -6.25, 106.7 -6.25, 106.7 -6.2)))', 4326))
ON CONFLICT (id) DO NOTHING;

-- Insert sample kelurahan data
INSERT INTO kelurahan (id, name, kecamatan_id, geom) VALUES 
(3201010001, 'Nanggung', 3201010, ST_GeomFromText('MULTIPOLYGON(((106.5 -6.3, 106.55 -6.3, 106.55 -6.35, 106.5 -6.35, 106.5 -6.3)))', 4326)),
(3201010002, 'Kalibunder', 3201010, ST_GeomFromText('MULTIPOLYGON(((106.55 -6.35, 106.6 -6.35, 106.6 -6.4, 106.55 -6.4, 106.55 -6.35)))', 4326)),
(3201020001, 'Leuwiliang', 3201020, ST_GeomFromText('MULTIPOLYGON(((106.6 -6.4, 106.65 -6.4, 106.65 -6.45, 106.6 -6.45, 106.6 -6.4)))', 4326)),
(3171010001, 'Jagakarsa', 3171010, ST_GeomFromText('MULTIPOLYGON(((106.7 -6.2, 106.73 -6.2, 106.73 -6.23, 106.7 -6.23, 106.7 -6.2)))', 4326))
ON CONFLICT (id) DO NOTHING;