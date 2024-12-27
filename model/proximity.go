package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Proximity Proximity
type Proximity struct {
	basic.Base
	basic.DataOwner
	ProximityAPI
	ProximityTranslation *ProximityTranslation `json:"proximity_translation,omitempty"` // Proximity Translation
}

// ProximityAPI Proximity API
type ProximityAPI struct {
	ProximityCode *int    `json:"proximity_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`          // Proximity Code
	ProximityName *string `json:"proximity_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Onsite"` // Proximity Name
}

// Seed data
func (s Proximity) Seed() *[]Proximity {
	data := []Proximity{}
	items := []string{
		"Onsite",
		"Offsite",
		"Nearby",
		"Information not available",
		"Onsite and offsite",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, Proximity{
			ProximityAPI: ProximityAPI{
				ProximityCode: &code,
				ProximityName: &name,
			},
		})
	}

	return &data
}
