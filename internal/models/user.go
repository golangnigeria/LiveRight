// Package models defines the core data structures for the LiveRight healthcare platform.
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

package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Role represents a user role (e.g., admin, doctor, patient)
type Role struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Email is a case-insensitive email type backed by PostgreSQL CITEXT.
// It behaves like a string but ensures consistent handling.
type Email string

// Value implements driver.Valuer for storing in DB (as TEXT/CITEXT)
func (e Email) Value() (driver.Value, error) {
	return string(e), nil
}

// Scan implements sql.Scanner for reading from DB
func (e *Email) Scan(value any) error {
	if value == nil {
		*e = ""
		return nil
	}
	if s, ok := value.(string); ok {
		*e = Email(s)
		return nil
	}
	return fmt.Errorf("cannot scan %T into Email", value)
}

// User represents a system user (patient, doctor, admin, etc.)
type User struct {
	ID           int64     `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	Email        Email     `json:"email" db:"email"`     // CITEXT → handled as string
	PasswordHash []byte    `json:"-" db:"password_hash"` // never expose in JSON
	RoleID       Role      `json:"role_id" db:"role_id"`
	Phone        *string   `json:"phone,omitempty" db:"phone"` // nullable
	Active       bool      `json:"active" db:"active"`
}

// HashPassword sets PasswordHash using bcrypt
func (u *User) HashPassword(plain string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = hashed
	return nil
}

// PasswordMatches compares password with hash
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(plainText))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
