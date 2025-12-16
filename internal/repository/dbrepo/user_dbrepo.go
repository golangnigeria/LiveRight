// Package dbrepo defines the core data structures for the LiveRight healthcare platform.
// These types map directly to the PostgreSQL database schema and are used across
// the application for data persistence, business logic, and API serialization.
//
// The package includes entities such as User, Role, Appointment, Wallet, LabTest,
// Pharmacy, and InsuranceClaim—representing the MVP feature set for patient,
// doctor, and administrative workflows.
//
// Special attention is given to PostgreSQL-specific types:
//   - Email uses CITEXT for case-insensitive storage and comparison.
//   - PasswordHash is stored as BYTEA and never exposed in JSON responses.
//   - Nullable fields (e.g., phone, notes) use pointers or sql.Null* types for safe handling.
//
// This package is database-agnostic in structure but designed to work seamlessly
// with standard Go SQL drivers (e.g., pq, pgx) and query libraries like sqlx or GORM.
package dbrepo

import (
	"database/sql"

	"github.com/golangnigeria/liveright_backend/internal/models"
)

func (m *PostgresDBRepo) InsertUser(user *models.User) (*models.User, error) {
	query := `
    INSERT INTO users (first_name, last_name, email, password_hash, role_id, phone, active)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, created_at
`

	var inserted models.User
	var phone sql.NullString
	if user.Phone != nil {
		phone = sql.NullString{String: *user.Phone, Valid: true}
	}

	err := m.DB.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.RoleID.ID,
		phone,
		user.Active,
	).Scan(&inserted.ID, &inserted.CreatedAt)
	if err != nil {
		return nil, err
	}

	// Fill the rest from input
	inserted.FirstName = user.FirstName
	inserted.LastName = user.LastName
	inserted.Email = user.Email
	inserted.PasswordHash = user.PasswordHash
	inserted.RoleID = user.RoleID
	inserted.Active = user.Active
	if phone.Valid {
		inserted.Phone = &phone.String
	}
	return &inserted, nil
}
