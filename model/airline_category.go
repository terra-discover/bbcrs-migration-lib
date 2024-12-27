package model

import (
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirlineCategory Airline Category
type AirlineCategory struct {
	basic.Base
	basic.DataOwner
	AirlineCategoryAPI
	AirlineCategoryTranslation *AirlineCategoryTranslation `json:"airline_category_translation,omitempty"`
	AirlineCategoryAirline     []AirlineCategoryAirline    `json:"airline_category_airline,omitempty"`
}

// AirlineCategoryAPI Airline Category API
type AirlineCategoryAPI struct {
	AirlineCategoryCode *string `json:"airline_category_code,omitempty" gorm:"type:varchar(2);not null;index:idx_airline_category_code_deleted_at,unique,where:deleted_at is null" example:"01"`                    // Airline Category Code
	AirlineCategoryName *string `json:"airline_category_name,omitempty" gorm:"type:varchar(64);not null;index:idx_airline_category_name_deleted_at,unique,where:deleted_at is null" example:"full service carrier"` // Airline Category Name
}

// Seed Seed data
func (f *AirlineCategory) Seed() *[]AirlineCategory {
	data := []AirlineCategory{}

	items := []string{
		"FS|Full Service Carrier",
		"LC|Low Cost Carrier",
	}

	for i := range items {
		content := strings.Split(items[i], "|")
		code := content[0]
		name := content[1]
		data = append(data, AirlineCategory{
			AirlineCategoryAPI: AirlineCategoryAPI{
				AirlineCategoryCode: &code,
				AirlineCategoryName: &name,
			},
		})
	}

	return &data
}
