package models

import (
	"time"
	"github.com/gofrs/uuid"
)

type Transaction struct {
	ID        int       `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	FarmID    *int      `json:"farm_id,omitempty" db:"farm_id"`       // Optional
	VendorID  *int      `json:"vendor_id,omitempty" db:"vendor_id"`   // Optional
	Type      string    `json:"type" db:"type"`
	Reference string    `json:"reference" db:"reference"`
	Amount    float64   `json:"amount" db:"amount"`
	Status    string    `json:"status" db:"status"`
	Meta      string    `json:"meta" db:"meta"` // Optional JSON string
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
