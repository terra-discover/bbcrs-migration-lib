package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RateTimeUnit Rate Time Unit
type RateTimeUnit struct {
	basic.Base
	basic.DataOwner
	RateTimeUnitAPI
	RateTimeUnitTranslation *RateTimeUnitTranslation `json:"rate_time_unit_translation,omitempty"` // Rate Time Unit Translation
}

// RateTimeUnitAPI Rate Time Unit API
type RateTimeUnitAPI struct {
	RateTimeUnitCode *int    `json:"rate_time_unit_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`        // Rate Time Unit Code
	RateTimeUnitName *string `json:"rate_time_unit_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Year"` // Rate Time Unit Name
}

// Seed data
func (s RateTimeUnit) Seed() *[]RateTimeUnit {
	data := []RateTimeUnit{}
	items := []string{
		"Year",
		"Month",
		"Week",
		"Day",
		"Hour",
		"Minute",
		"Second",
		"Full Duration",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, RateTimeUnit{
			RateTimeUnitAPI: RateTimeUnitAPI{
				RateTimeUnitCode: &code,
				RateTimeUnitName: &name,
			},
		})
	}

	return &data
}
