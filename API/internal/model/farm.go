package models

import (
	"fmt"
	"time"

	"github.com/golangnigeria/MyNneFarm/internal/lib/date"
	"github.com/golangnigeria/MyNneFarm/internal/lib/float"
)

// Farm represents a farm available for investment
type FarmModel struct {
	ID                 int           `json:"id" db:"id"`                                   // Auto-increment primary key
	Title              string        `json:"title" db:"title"`                             // e.g. "Oyo Maize Farm"
	Crop               string        `json:"crop" db:"crop"`                               // e.g. "Maize", "Cassava"
	Location           string        `json:"location" db:"location"`                       // e.g. "Oyo, Nigeria"
	ImageURL           string        `json:"image_url" db:"image_url"`                     // Cloud image for display
	Description        string        `json:"description" db:"description"`                 // Full description for display
	PricePerUnit       float.Price   `json:"price_per_unit" db:"price_per_unit"`           // e.g. 5000.00 NGN
	ExpectedROI        float.Price   `json:"expected_roi" db:"expected_roi"`               // e.g. 1.2 = 20% profit
	ExpectedYield      float.Price   `json:"expected_yield" db:"expected_yield"`           // e.g. 10000 kg
	ExpectedRevenue    float.Price   `json:"expected_revenue" db:"expected_revenue"`       // e.g. 6000000.00 NGN
	ProductionDuration int           `json:"production_duration" db:"production_duration"` // e.g. 90 days
	UnitsAvailable     int           `json:"units_available" db:"units_available"`
	UnitsSold          int           `json:"units_sold" db:"units_sold"`
	StartDate          date.DateOnly `json:"start_date" db:"start_date"`     // Date the farm starts
	HarvestDate        date.DateOnly `json:"harvest_date" db:"harvest_date"` // Date users receive Food Credits
	IsActive           bool          `json:"is_active" db:"is_active"`       // Whether it's open for investment
	CreatedAt          time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at" db:"updated_at"`
}

// Validate checks if the farm data is valid
func (f *FarmModel) Validate() error {
	if f.Title == "" {
		return fmt.Errorf("title is required")
	}
	if f.Crop == "" {
		return fmt.Errorf("crop is required")
	}
	if f.Location == "" {
		return fmt.Errorf("location is required")
	}
	if f.PricePerUnit <= 0 {
		return fmt.Errorf("price per unit must be greater than zero")
	}
	if f.ExpectedROI <= 0 {
		return fmt.Errorf("expected ROI must be greater than zero")
	}
	if f.ExpectedYield <= 0 {
		return fmt.Errorf("expected yield must be greater than zero")
	}
	if f.ProductionDuration <= 0 {
		return fmt.Errorf("production duration must be greater than zero")
	}
	 
	return nil
}
