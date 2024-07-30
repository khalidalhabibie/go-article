

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Asia/Jakarta";

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    address TEXT,
    phone_no VARCHAR(50),
    password VARCHAR(255),
    verified_at TIMESTAMP WITHOUT TIME ZONE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for the Users Table
CREATE INDEX idx_users_name ON users (name);
CREATE INDEX idx_users_email ON users (email);
