package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Region Region
type Region struct {
	basic.Base
	basic.DataOwner
	RegionAPI
	RegionTranslation *RegionTranslation `json:"region_translation,omitempty"`
}

// RegionAPI Region API
type RegionAPI struct {
	RegionCode *string `json:"region_code,omitempty" gorm:"type:varchar(2);not null" example:"IN"`          // Region Code
	RegionName *string `json:"region_name,omitempty" gorm:"type:varchar(256);not null" example:"Indonesia"` // Region Name
}
