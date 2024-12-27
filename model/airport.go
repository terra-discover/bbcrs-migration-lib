package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Airport Airport
type Airport struct {
	basic.Base
	basic.DataOwner
	AirportAPI
	City               *City               `json:"city,omitempty"` // City
	CityCode           *string             `json:"city_code,omitempty" gorm:"-"`
	CityName           *string             `json:"city_name,omitempty" gorm:"-"`
	CountryName        *string             `json:"country_name,omitempty" gorm:"-"`
	AirportTranslation *AirportTranslation `json:"airport_translation,omitempty"` // Airport Translation
}

// AirportAPI Airport API
// @issue https://trello.com/c/LRimu57N
type AirportAPI struct {
	AirportCode *string    `json:"airport_code,omitempty" gorm:"type:varchar(3);uniqueIndex:,where:deleted_at is null;not null" example:"CGK"`                 // 3 Characters airport_code
	AirportName *string    `json:"airport_name,omitempty" gorm:"type:varchar(256);not null" example:"Bandara Soekarno Hatta"`                                  // Airport Name
	CityID      *uuid.UUID `json:"city_id,omitempty" swaggertype:"string" format:"uuid"`                                                                       // City ID
	IcaoCode    *string    `json:"icao_code,omitempty" gorm:"type:varchar(4);uniqueIndex:,where:deleted_at is null and icao_code is not null;" example:"WIII"` // Icao Code
}
