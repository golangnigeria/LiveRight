package repo

import (
	"database/sql"
	"time"
)

type NeonDBRepo struct {
	DB *sql.DB
}

const dbTimeout = 5 * time.Second

func (m *NeonDBRepo) Connection() *sql.DB {
	return m.DB
}
