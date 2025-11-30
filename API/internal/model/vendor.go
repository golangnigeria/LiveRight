package models

import "time"

type Vendor struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Email         string    `json:"email" db:"email"`
	Phone         string    `json:"phone" db:"phone"`
	Location      string    `json:"location" db:"location"`
	Verified      bool      `json:"verified" db:"verified"`
	LogoURL       string    `json:"logo_url" db:"logo_url"`
	WalletBalance float64   `json:"wallet_balance" db:"wallet_balance"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
