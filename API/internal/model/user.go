package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int            `json:"id" db:"id"`
	FullName      string         `json:"full_name" db:"full_name"`
	Email         string         `json:"email" db:"email"`
	Phone         string         `json:"phone" db:"phone"`
	Password      Password       `json:"-" db:"password"` // stored as hash, omitted in JSON
	Roles         pq.StringArray `json:"roles,omitempty" db:"roles"`
	WalletBalance int            `json:"wallet_balance" db:"wallet_balance"`
	Activated     bool           `json:"activated" db:"activated"`
	Version       int            `json:"version" db:"version"`
	ReferredBy    string         `json:"referred_by,omitempty" db:"referred_by"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
}

// Password handles both the plain and hashed password representation
type Password struct {
	Plain *string
	Hash  []byte
}

// Set hashes the plain-text password and stores it in Hash
func (p *Password) Set(plainTextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 12)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	p.Plain = &plainTextPassword
	p.Hash = hash
	return nil
}

// Match compares a plain-text password with the stored hash
func (p *Password) Match(plainTextPassword string) (bool, error) {

	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plainTextPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, fmt.Errorf("password comparison failed: %w", err)
	}
	return true, nil
}

// PasswordMatches checks if the given plain-text password matches the stored hash
func (u *User) PasswordMatches(plainText string) (bool, error) {
	return u.Password.Match(plainText)
}
