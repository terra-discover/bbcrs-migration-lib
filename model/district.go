package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// District District
type District struct {
	basic.Base
	basic.DataOwner
	DistrictAPI
	DistrictTranslation *DistrictTranslation `json:"district_translation,omitempty"` // District Translation
	AdministrativeCity  *AdministrativeCity  `json:"administrative_city,omitempty"`  // Administrative City
	City                *City                `json:"city,omitempty"`                 // City
}

// DistrictAPI District API
type DistrictAPI struct {
	DistrictCode         *string    `json:"district_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;index:idx_district__district_name,unique,where:deleted_at is null;not null" example:"Gambir"` // District Code
	DistrictName         *string    `json:"district_name,omitempty" gorm:"type:varchar(256);index:idx_district__district_name,unique,where:deleted_at is null;not null" example:"Gambir"`                                       // District Name
	AdministrativeCityID *uuid.UUID `json:"administrative_city_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                                                // Administrative City ID
	CityID               *uuid.UUID `json:"city_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                                                               // City ID
}
