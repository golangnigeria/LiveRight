// Package dbrepo provides concrete implementations of the repository
// interfaces using PostgreSQL as the data store. It contains all database-
// level operations, SQL queries, and persistence logic required by the
// application.
package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golangnigeria/liveright_backend/internal/models"
)

// PostgresDBRepo is the PostgreSQL implementation of the application's
// DatabaseRepo interface. It holds a database connection pool and exposes
// methods for executing queries against the PostgreSQL backend.
type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := ` 
				SELECT id, created_at, first_name, last_name, email, password_hash, role_id,
				phone, active from users where email = $1
	`

	var user models.User
	var roleID int64
	var phone sql.NullString

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&roleID,
		&phone,
		&user.Active,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	user.RoleID = models.Role{ID: roleID}
	if phone.Valid {
		user.Phone = &phone.String
	}

	return &user, nil
}
