-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create countries table
CREATE TABLE IF NOT EXISTS countries (
                                         id SERIAL PRIMARY KEY,
                                         name VARCHAR(255) NOT NULL UNIQUE,
    country_code VARCHAR(2) NOT NULL UNIQUE,
    calling_code VARCHAR(10) NOT NULL,
    flag_url VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
    );

CREATE INDEX idx_countries_deleted_at ON countries(deleted_at);

-- Create clients table
CREATE TABLE IF NOT EXISTS clients (
                                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255) NOT NULL,
    middle_name VARCHAR(255),
    last_name VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    pesel VARCHAR(50) NOT NULL UNIQUE,
    phone_number VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    passcode VARCHAR(10),
    sub UUID,
    registration_status VARCHAR(50) NOT NULL DEFAULT 'STARTED',
    version INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
    );

CREATE INDEX idx_clients_pesel ON clients(pesel);
CREATE INDEX idx_clients_email ON clients(email);
CREATE INDEX idx_clients_phone_number ON clients(phone_number);
CREATE INDEX idx_clients_deleted_at ON clients(deleted_at);

-- Create addresses table
CREATE TABLE IF NOT EXISTS addresses (
                                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    address_type VARCHAR(20) NOT NULL,
    postcode VARCHAR(20),
    street_name VARCHAR(255) NOT NULL,
    building_number VARCHAR(50) NOT NULL,
    apartment_number VARCHAR(50),
    city VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
    );

CREATE INDEX idx_addresses_user_id ON addresses(user_id);
CREATE INDEX idx_addresses_deleted_at ON addresses(deleted_at);

-- Create documents table
CREATE TABLE IF NOT EXISTS documents (
                                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    document_type VARCHAR(30) NOT NULL,
    document_number VARCHAR(100) NOT NULL UNIQUE,
    issuing_country VARCHAR(100) NOT NULL,
    issue_date DATE NOT NULL,
    expiry_date DATE NOT NULL,
    notification_sent BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT chk_issue_date CHECK (issue_date >= '1900-01-01'),
    CONSTRAINT chk_expiry_after_issue CHECK (expiry_date > issue_date)
    );

CREATE INDEX idx_documents_user_id ON documents(user_id);
CREATE INDEX idx_documents_document_number ON documents(document_number);
CREATE INDEX idx_documents_deleted_at ON documents(deleted_at);

-- Create photos table
CREATE TABLE IF NOT EXISTS photos (
                                      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    url VARCHAR(500) NOT NULL,
    photo_type VARCHAR(20) NOT NULL,
    original_filename VARCHAR(255),
    content_type VARCHAR(100),
    file_size BIGINT,
    client_id UUID REFERENCES clients(id) ON DELETE CASCADE,
    document_id UUID REFERENCES documents(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT chk_photo_relation CHECK (
(client_id IS NOT NULL AND document_id IS NULL) OR
(client_id IS NULL AND document_id IS NOT NULL)
    )
    );

CREATE INDEX idx_photos_client_id ON photos(client_id);
CREATE INDEX idx_photos_document_id ON photos(document_id);
CREATE INDEX idx_photos_deleted_at ON photos(deleted_at);

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply triggers to all tables
CREATE TRIGGER update_countries_updated_at BEFORE UPDATE ON countries
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_clients_updated_at BEFORE UPDATE ON clients
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_addresses_updated_at BEFORE UPDATE ON addresses
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_documents_updated_at BEFORE UPDATE ON documents
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_photos_updated_at BEFORE UPDATE ON photos
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();