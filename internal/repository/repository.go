// Package repository defines the interfaces and contracts for data access
// across the application. It abstracts database operations behind reusable
// repository patterns, enabling clean separation between business logic
// and persistence layers.
package repository

import (
	"database/sql"

	"github.com/golangnigeria/liveright_backend/internal/models"
)

// DatabaseRepo describes the set of methods required for interacting with
// the application's data storage layer. Concrete implementations may use
// PostgreSQL, MongoDB, or any other database engine.
type DatabaseRepo interface {
	Connection() *sql.DB
	AllDoctors() ([]*models.Doctor, error)
}
