package repo

import (
	"context"
	"fmt"
	"time"

	models "github.com/golangnigeria/MyNneFarm/internal/model"
)



// GetFarms retrieves all farms from the database.
func (m *NeonDBRepo) GetFarms() ([]models.FarmModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := ` SELECT 
					id, title, crop, location, coalesce(image_url, '') AS image_url, description, price_per_unit, expected_roi, units_available, units_sold, start_date, harvest_date, is_active, created_at, updated_at 
				FROM 
				  	farms
				ORDER BY created_at DESC;`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var farms []models.FarmModel

	for rows.Next() {
		var farm models.FarmModel
		if err := rows.Scan(
			&farm.ID,
			&farm.Title,
			&farm.Crop,
			&farm.Location,
			&farm.ImageURL,
			&farm.Description,
			&farm.PricePerUnit,
			&farm.ExpectedROI,
			&farm.UnitsAvailable,
			&farm.UnitsSold,
			&farm.StartDate,
			&farm.HarvestDate,
			&farm.IsActive,
			&farm.CreatedAt,
			&farm.UpdatedAt); err != nil {
			return nil, err
		}
		farms = append(farms, farm)
	}

	return farms, nil
}

// InsertFarm inserts a new farm into the database.
func (m *NeonDBRepo) InsertFarm(farm models.FarmModel) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO farms (title, crop, location, image_url, description, price_per_unit, expected_roi, expected_yield, expected_revenue, production_duration, units_available, start_date, harvest_date, is_active, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, NOW(), NOW()) 
			  RETURNING id;`

	err := m.DB.QueryRowContext(ctx,
		query,
		farm.Title,
		farm.Crop,
		farm.Location,
		farm.ImageURL,
		farm.Description,
		farm.PricePerUnit,
		farm.ExpectedROI,
		farm.ExpectedYield,
		farm.ExpectedRevenue,
		farm.ProductionDuration,
		farm.UnitsAvailable,
		farm.StartDate.ToTime(),
		farm.HarvestDate.ToTime(),
		farm.IsActive).Scan(&farm.ID)

	if err != nil {
		return 0, err
	}

	// Update the CreatedAt and UpdatedAt fields
	farm.CreatedAt = time.Now()
	farm.UpdatedAt = time.Now()

	// Return the ID of the newly inserted farm
	if farm.ID <= 0 {
		return 0, fmt.Errorf("failed to insert farm")
	}

	// convert farm id int to int64
	return int64(farm.ID), nil
}
