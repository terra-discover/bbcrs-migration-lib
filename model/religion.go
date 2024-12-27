package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// Religion Religion
type Religion struct {
	basic.Base
	basic.DataOwner
	ReligionAPI
	ReligionTranslation *ReligionTranslation `json:"religion_translation,omitempty"`
}

// ReligionAPI Religion API
type ReligionAPI struct {
	ReligionCode *string `json:"religion_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`  //Religion Code
	ReligionName *string `json:"religion_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Religion Name
}
