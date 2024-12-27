package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// StateProvince State Province
type StateProvince struct {
	basic.Base
	basic.DataOwner
	StateProvinceAPI
	StateProvinceTranslation *StateProvinceTranslation `json:"state_province_translation,omitempty"` // state province translation
	Country                  *Country                  `json:"country,omitempty"`                    // country
	StateProvinceCategory    *StateProvinceCategory    `json:"state_province_category,omitempty"`    // state province category
}

// StateProvinceAPI State Province API
type StateProvinceAPI struct {
	StateProvinceCode       *string    `json:"state_province_code,omitempty" gorm:"type:varchar(8);index:idx_state_province_state_province_code,unique,where:deleted_at is null;not null" example:"JOG"`                                                                                  // State Province Code
	StateProvinceName       *string    `json:"state_province_name,omitempty" gorm:"type:varchar(256);index:idx_state_province_state_province_name,unique,where:deleted_at is null;not null" example:"Jogjakarta"`                                                                         // State Province Name
	StateProvinceCategoryID *uuid.UUID `json:"state_province_category_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                                                                                                           // State Province Category ID
	CountryID               *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36);index:idx_state_province_state_province_code,unique,where:deleted_at is null;index:idx_state_province_state_province_name,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"` // Country ID
}
