package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
	"github.com/lib/pq"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrUserNotFound   = errors.New("user not found")
)

func (m *NeonDBRepo) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Normalize inputs
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.Phone = strings.TrimSpace(user.Phone)
	user.FullName = strings.TrimSpace(user.FullName)

	if len(user.Roles) == 0 {
		user.Roles = []string{"farmer"} // default role
	}

	now := time.Now().UTC()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	query := `
		INSERT INTO users (
			full_name, email, phone, password, roles,
			wallet_balance, activated, version, referred_by,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11
		) RETURNING id, version, created_at
	`

	err := m.DB.QueryRowContext(ctx, query,
		user.FullName,
		user.Email,
		user.Phone,
		user.Password.Hash,
		pq.Array(user.Roles),
		user.WalletBalance,
		user.Activated,
		user.Version,
		user.ReferredBy,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID, &user.Version, &user.CreatedAt)

	if err != nil {
		if strings.Contains(err.Error(), `users_email_key`) {
			return ErrDuplicateEmail
		}

		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (m *NeonDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, full_name, email, phone, password, roles,
		       wallet_balance, activated, version, referred_by,
		       created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Phone,
		&user.Password.Hash,
		&user.Roles,
		&user.WalletBalance,
		&user.Activated,
		&user.Version,
		&user.ReferredBy,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("querying user by email: %w", err)
	}

	return &user, nil
}


func (m *NeonDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, full_name, email, phone, password, roles,
		       wallet_balance, activated, version, referred_by,
		       created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Phone,
		&user.Password.Hash,
		&user.Roles,
		&user.WalletBalance,
		&user.Activated,
		&user.Version,
		&user.ReferredBy,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("querying user by email: %w", err)
	}

	return &user, nil
}

// UpdateUser updates an existing user in the database.
func (m *NeonDBRepo) UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
			UPDATE users
				SET full_name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
			WHERE id = $5 AND version = $6
			RETURNING version`

	err := m.DB.QueryRowContext(ctx, query,
		user.FullName,
		user.Email,
		user.Password,
		user.Activated,
		user.ID,
		user.Version,
	).Scan(&user.Version)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			return ErrUserNotFound
		default:
			return err
		}
	}

	return nil
}
