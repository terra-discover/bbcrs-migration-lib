package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// MaritalStatus Marital Status
type MaritalStatus struct {
	basic.Base
	basic.DataOwner
	MaritalStatusAPI
	MaritalStatusTranslation *MaritalStatusTranslation `json:"marital_status_translation,omitempty"` // Marital Status Translation
}

// MaritalStatusAPI Marital Status API
type MaritalStatusAPI struct {
	MaritalStatusCode *string `json:"marital_status_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" example:"mrd"`      // Marital Status Code
	MaritalStatusName *string `json:"marital_status_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"married"` // Marital Status Name
}
