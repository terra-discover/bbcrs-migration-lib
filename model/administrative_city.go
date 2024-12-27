package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdministrativeCity Administrative City
type AdministrativeCity struct {
	basic.Base
	basic.DataOwner
	AdministrativeCityAPI
	AdministrativeCityTranslation *AdministrativeCityTranslation `json:"administrative_city_translation,omitempty"` // Administrative City Translation
	CapitalCity                   *City                          `json:"capital_city,omitempty"`
}

// AdministrativeCityAPI Administrative City API
type AdministrativeCityAPI struct {
	CityCode        *string    `json:"city_code,omitempty" gorm:"type:varchar(3);index:idx_administrative_city__city_name,unique,where:deleted_at is null;not null" example:"JKT"`               // City Code
	CityName        *string    `json:"city_name,omitempty" gorm:"type:varchar(256);index:idx_administrative_city__city_name,unique,where:deleted_at is null;not null" example:"South Jakarta"`   // City Name
	CountryID       *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36);index:idx_administrative_city__city_name,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"` // Country ID
	StateProvinceID *uuid.UUID `json:"state_province_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                                   // State Province ID
	CapitalCityID   *uuid.UUID `json:"capital_city_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                                     // Capital City ID
}
