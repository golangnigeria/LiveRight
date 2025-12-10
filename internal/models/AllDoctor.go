// Package models defines the core data structures used throughout the
// application. These models represent the entities stored in the database
// and exposed through the API layer.
package models

import "time"

// Doctor represents a medical professional with personal details,
// specialization info, and profile metadata used across the application.
type Doctor struct {
	ID                int64     `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	Specialization    string    `json:"specialization"`
	YearsOfExperience int       `json:"years_of_experience"`
	Bio               string    `json:"bio"`
	ProfileImage      string    `json:"profile_image"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}
