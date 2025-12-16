-- +goose Up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Seed roles
INSERT INTO roles (id, name, created_at) VALUES (1, 'patient', now()) ON CONFLICT DO NOTHING;
INSERT INTO roles (id, name, created_at) VALUES (2, 'doctor', now()) ON CONFLICT DO NOTHING;
INSERT INTO roles (id, name, created_at) VALUES (3, 'lab', now()) ON CONFLICT DO NOTHING;
INSERT INTO roles (id, name, created_at) VALUES (4, 'pharmacy', now()) ON CONFLICT DO NOTHING;
INSERT INTO roles (id, name, created_at) VALUES (5, 'insurance', now()) ON CONFLICT DO NOTHING;
INSERT INTO roles (id, name, created_at) VALUES (6, 'admin', now()) ON CONFLICT DO NOTHING;

-- +goose Down
DELETE FROM roles WHERE id IN (1,2,3,4,5,6);
DROP TABLE IF EXISTS roles;
DROP EXTENSION IF EXISTS citext;

