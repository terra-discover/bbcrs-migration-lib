package model

import (
	"strings"

	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// OffsetTimeUnit Offset Time Unit
type OffsetTimeUnit struct {
	basic.Base
	basic.DataOwner
	OffsetTimeUnitAPI
	OffsetTimeUnitTranslation *OffsetTimeUnitTranslation `json:"offset_time_unit_translation,omitempty"` // Offset Time Unit Translation
}

// OffsetTimeUnitAPI Offset Time Unit API
type OffsetTimeUnitAPI struct {
	OffsetTimeUnitCode *string `json:"offset_time_unit_code,omitempty" gorm:"type:varchar(16);index:,unique,where:deleted_at is null;not null" example:"Year"`  // Offset Time Unit Code
	OffsetTimeUnitName *string `json:"offset_time_unit_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Year"` // Offset Time Unit Name
}

// Seed data
func (s OffsetTimeUnit) Seed() *[]OffsetTimeUnit {
	data := []OffsetTimeUnit{}
	items := []string{
		"Year|Year",
		"Month|Month",
		"Week|Week",
		"Day|Day",
		"Hour|Hour",
	}

	for i := range items {
		contents := strings.Split(items[i], "|")
		var code string = contents[0]
		var name string = contents[1]
		data = append(data, OffsetTimeUnit{
			OffsetTimeUnitAPI: OffsetTimeUnitAPI{
				OffsetTimeUnitCode: &code,
				OffsetTimeUnitName: &name,
			},
		})
	}

	return &data
}
