DROP TRIGGER IF EXISTS update_photos_updated_at ON photos;
DROP TRIGGER IF EXISTS update_documents_updated_at ON documents;
DROP TRIGGER IF EXISTS update_addresses_updated_at ON addresses;
DROP TRIGGER IF EXISTS update_clients_updated_at ON clients;
DROP TRIGGER IF EXISTS update_countires_updated_at ON countries;

DROP FUNCTION IF EXISTS update_updated_at_column();

DROP TABLE IF EXISTS photos;
DROP TABLE IF EXISTS documents;
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS countries;


DROP EXTENSION IF EXISTS "uuid-ossp"