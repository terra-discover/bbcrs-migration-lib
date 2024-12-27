package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// TransportationMode Transportation Mode
type TransportationMode struct {
	basic.Base
	basic.DataOwner
	TransportationModeAPI
	TransportationModeTranslation *TransportationModeTranslation `json:"transportation_mode_translation,omitempty"` // Transportation Mode Translation
}

// TransportationModeAPI Transportation Mode API
type TransportationModeAPI struct {
	TransportationModeCode *int    `json:"transportation_mode_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`           // Transportation Mode Code
	TransportationModeName *string `json:"transportation_mode_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"Bicycle"` // Transportation Mode Name
}

// Seed data
func (r TransportationMode) Seed() *[]TransportationMode {
	data := []TransportationMode{}
	items := []string{
		"Bicycle",
		"Boat",
		"Bus",
		"Cable car",
		"Car",
		"Carriage",
		"Courtesy car",
		"Helicopter",
		"Limousine",
		"Metro",
		"Monorail",
		"Motorbike",
		"Pack animal",
		"Plane",
		"Rental car",
		"Rickshaw",
		"Shuttle",
		"Subway",
		"Sedan chair",
		"Taxi",
		"Train",
		"Trolley",
		"Tube",
		"Walk",
		"Water taxi",
		"Other or alternate",
		"Car/rush hour",
		"Taxi/rush hour",
		"No transportation",
		"Express train",
		"Public",
		"Alternate",
		"Ferry",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, TransportationMode{
			TransportationModeAPI: TransportationModeAPI{
				TransportationModeCode: &code,
				TransportationModeName: &name,
			},
		})
	}

	return &data
}
