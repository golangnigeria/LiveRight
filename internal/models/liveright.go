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
	"time"
)

// Appointment between patient and doctor
type Appointment struct {
	ID              int64     `json:"id" db:"id"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	PatientID       int64     `json:"patient_id" db:"patient_id"`
	DoctorID        int64     `json:"doctor_id" db:"doctor_id"`
	AppointmentTime time.Time `json:"appointment_time" db:"appointment_time"`
	Status          string    `json:"status" db:"status"` // e.g., pending, confirmed, completed
	Notes           *string   `json:"notes,omitempty" db:"notes"`
}

// LRCWallet (LiveRight Card Wallet)
type LRCWallet struct {
	ID            int64     `json:"id" db:"id"`
	UserID        int64     `json:"user_id" db:"user_id"`
	Balance       float64   `json:"balance" db:"balance"` // NUMERIC(12,2) → float64 is safe for money in MVP; consider decimal for production
	RewardsPoints int       `json:"rewards_points" db:"rewards_points"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// Transaction (topup, payment, refund)
type Transaction struct {
	ID        int64     `json:"id" db:"id"`
	WalletID  int64     `json:"wallet_id" db:"wallet_id"`
	Amount    float64   `json:"amount" db:"amount"`
	Type      string    `json:"type" db:"type"`     // "topup", "payment", "refund"
	Status    string    `json:"status" db:"status"` // "pending", "completed", etc.
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Lab facility
type Lab struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Address   *string   `json:"address,omitempty" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// LabTest ordered for a patient
type LabTest struct {
	ID        int64     `json:"id" db:"id"`
	LabID     int64     `json:"lab_id" db:"lab_id"`
	PatientID int64     `json:"patient_id" db:"patient_id"`
	TestName  string    `json:"test_name" db:"test_name"`
	Result    *string   `json:"result,omitempty" db:"result"`
	Status    string    `json:"status" db:"status"` // "pending", "completed", etc.
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Pharmacy for the list of approved onces
type Pharmacy struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Address   *string   `json:"address,omitempty" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Insurer (insurance provider)
type Insurer struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// InsuranceClaim for those Patient that have it
type InsuranceClaim struct {
	ID        int64     `json:"id" db:"id"`
	PatientID int64     `json:"patient_id" db:"patient_id"`
	InsurerID int64     `json:"insurer_id" db:"insurer_id"`
	Amount    float64   `json:"amount" db:"amount"`
	Status    string    `json:"status" db:"status"` // "pending", "approved", "rejected"
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
