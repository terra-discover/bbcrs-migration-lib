package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Address struct
type Address struct {
	basic.Base
	basic.DataOwner
	AddressAPI
	City          *City          `json:"city,omitempty"`           // City
	Country       *Country       `json:"country,omitempty"`        // Country
	StateProvince *StateProvince `json:"state_province,omitempty"` // State Province
	Destination   *Destination   `json:"destination,omitempty"`    // Destination
	Zone          *Zone          `json:"zone,omitempty"`           // Zone
	District      *District      `json:"district,omitempty"`       // District
	Subdistrict   *Subdistrict   `json:"subdistrict,omitempty"`    // Subdistrict
	//AdministrativeCity    *AdministrativeCityID
	//AltitudeUnitOfMeasure *AltitudeUnitOfMeasure
}

// AddressAPI struct
type AddressAPI struct {
	AddressLine             *string    `json:"address_line,omitempty" gorm:"type:varchar(512)" example:"street fighter"`                         // Address Line
	BuildingRoom            *string    `json:"building_room,omitempty" gorm:"type:varchar(256)" example:"no 4"`                                  // Building Room
	CityID                  *uuid.UUID `json:"city_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                     // City ID
	CityOther               *string    `json:"city_other,omitempty" gorm:"type:varchar(64)"`                                                     // City Other
	CountryID               *uuid.UUID `json:"country_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                  // Country ID
	County                  *string    `json:"county,omitempty" gorm:"type:varchar(32)"`                                                         // County
	PostalCode              *string    `json:"postal_code,omitempty" gorm:"type:varchar(16)" example:"12345"`                                    // Postal Code
	StateProvinceID         *uuid.UUID `json:"state_province_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`           // State Province ID
	StateProvinceOther      *string    `json:"state_province_other,omitempty" gorm:"type:varchar(128)"`                                          // State Province Other
	AdministrativeCityID    *uuid.UUID `json:"administrative_city_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`      // Administrative City ID
	DistrictID              *uuid.UUID `json:"district_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                 // District ID
	SubdistrictID           *uuid.UUID `json:"subdistrict_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`              // Subdistrict ID
	StreetNumber            *string    `json:"street_number,omitempty" gorm:"type:varchar(64)" example:"C4"`                                     // Street Number
	Unformatted             *bool      `json:"unformatted,omitempty"`                                                                            // Unformatted
	DestinationID           *uuid.UUID `json:"destination_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`              // Destination ID
	ZoneID                  *uuid.UUID `json:"zone_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"`                     // Zone ID
	Latitude                *float64   `json:"latitude,omitempty" example:"0"`                                                                   // Latitude
	Longitude               *float64   `json:"longitude,omitempty" example:"0"`                                                                  // Longitude
	Altitude                *float64   `json:"altitude,omitempty" example:"0"`                                                                   // Altitude
	AltitudeUnitOfMeasureID *uuid.UUID `json:"altitude_unit_of_measure_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36)"` // Altitude Unit Of Measure ID
}
