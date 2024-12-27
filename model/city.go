package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// City City
type City struct {
	basic.Base
	basic.DataOwner
	CityAPI
	CityTranslation    *CityTranslation     `json:"city_translation,omitempty"` // translations
	CountryName        *string              `json:"country_name,omitempty" gorm:"-"`
	Country            *Country             `json:"country,omitempty" gorm:"foreignKey:CountryID;references:ID"`              // country
	StateProvince      *StateProvince       `json:"state_province,omitempty" gorm:"foreignKey:StateProvinceID;references:ID"` // state province
	Airports           []Airport            `json:"airports,omitempty" gorm:"foreignKey:CityID;references:ID"`
	AdministrativeCity []AdministrativeCity `json:"administrative_city,omitempty" gorm:"foreignKey:CapitalCityID;references:ID"`
}

// CityAPI City API
type CityAPI struct {
	CityCode        *string    `json:"city_code,omitempty" gorm:"type:varchar(5);not null;index:idx_city_city_code,unique,where:deleted_at is null" example:"JKT"`                                                                               // City Code
	CountryID       *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36);index:idx_city_city_name,unique,where:deleted_at is null;index:idx_city_city_code,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`        // Country ID
	StateProvinceID *uuid.UUID `json:"state_province_id,omitempty" gorm:"type:varchar(36);index:idx_city_city_name,unique,where:deleted_at is null;index:idx_city_city_code,unique,where:deleted_at is null" format:"uuid" swaggertype:"string"` // State Province ID
	CityName        *string    `json:"city_name,omitempty" gorm:"type:varchar(256);not null;index:idx_city_city_name,unique,where:deleted_at is null" example:"Jakarta"`                                                                         // City Name
}
