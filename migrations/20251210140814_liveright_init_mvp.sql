-- +goose Up
-- ----------------------
-- Up Migration: Create production MVP schema for LiveRight
-- ----------------------

-- Roles
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Users
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    name TEXT NOT NULL,
    email CITEXT UNIQUE NOT NULL,
    password_hash BYTEA NOT NULL,
    role_id INT REFERENCES roles(id) ON DELETE SET NULL,
    phone TEXT,
    active BOOLEAN DEFAULT TRUE
);

-- Appointments
CREATE TABLE IF NOT EXISTS appointments (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    patient_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    doctor_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    appointment_time TIMESTAMP WITH TIME ZONE NOT NULL,
    status TEXT DEFAULT 'pending',
    notes TEXT
);

-- LiveRight Card Wallets
CREATE TABLE IF NOT EXISTS lrc_wallets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    balance NUMERIC(12,2) DEFAULT 0.00,
    rewards_points INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Transactions
CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    wallet_id BIGINT REFERENCES lrc_wallets(id) ON DELETE CASCADE,
    amount NUMERIC(12,2) NOT NULL,
    type TEXT NOT NULL, -- topup, payment, refund
    status TEXT DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Labs
CREATE TABLE IF NOT EXISTS labs (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Lab Tests
CREATE TABLE IF NOT EXISTS lab_tests (
    id BIGSERIAL PRIMARY KEY,
    lab_id BIGINT REFERENCES labs(id) ON DELETE CASCADE,
    patient_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    test_name TEXT NOT NULL,
    result TEXT,
    status TEXT DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Pharmacies
CREATE TABLE IF NOT EXISTS pharmacies (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Insurance Providers
CREATE TABLE IF NOT EXISTS insurers (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Insurance Claims
CREATE TABLE IF NOT EXISTS insurance_claims (
    id BIGSERIAL PRIMARY KEY,
    patient_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    insurer_id BIGINT REFERENCES insurers(id) ON DELETE CASCADE,
    amount NUMERIC(12,2) NOT NULL,
    status TEXT DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_appointments_patient_id ON appointments(patient_id);
CREATE INDEX IF NOT EXISTS idx_transactions_wallet_id ON transactions(wallet_id);
CREATE INDEX IF NOT EXISTS idx_lab_tests_patient_id ON lab_tests(patient_id);
CREATE INDEX IF NOT EXISTS idx_insurance_claims_patient_id ON insurance_claims(patient_id);

-- +goose Down
-- ----------------------
-- Down Migration: Drop all tables in reverse order
-- ----------------------

DROP INDEX IF EXISTS idx_insurance_claims_patient_id;
DROP INDEX IF EXISTS idx_lab_tests_patient_id;
DROP INDEX IF EXISTS idx_transactions_wallet_id;
DROP INDEX IF EXISTS idx_appointments_patient_id;
DROP INDEX IF EXISTS idx_users_email;

DROP TABLE IF EXISTS insurance_claims;
DROP TABLE IF EXISTS insurers;
DROP TABLE IF EXISTS pharmacies;
DROP TABLE IF EXISTS lab_tests;
DROP TABLE IF EXISTS labs;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS lrc_wallets;
DROP TABLE IF EXISTS appointments;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;

