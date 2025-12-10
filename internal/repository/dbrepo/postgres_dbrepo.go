// Package dbrepo provides concrete implementations of the repository
// interfaces using PostgreSQL as the data store. It contains all database-
// level operations, SQL queries, and persistence logic required by the
// application.
package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
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

// AllDoctors returns a list of all doctors stored in the database.
// It executes a SELECT query and maps the result to the Doctor model.
func (m *PostgresDBRepo) AllDoctors() ([]*models.Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        SELECT
            id,
            first_name,
            last_name,
            email,
            phone,
            specialization,
            years_of_experience,
            bio,
            coalesce(profile_image,''),
            created_at,
            updated_at
        FROM doctors
        ORDER BY id;
    `

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows", err)
		}
	}()

	var doctors []*models.Doctor

	for rows.Next() {
		var d models.Doctor

		err := rows.Scan(
			&d.ID,
			&d.FirstName,
			&d.LastName,
			&d.Email,
			&d.Phone,
			&d.Specialization,
			&d.YearsOfExperience,
			&d.Bio,
			&d.ProfileImage,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		doctors = append(doctors, &d)
	}

	return doctors, nil
}
