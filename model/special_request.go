package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// SpecialRequest Special Request
type SpecialRequest struct {
	basic.Base
	basic.DataOwner
	SpecialRequestAPI
	SpecialRequestTranslation *SpecialRequestTranslation `json:"special_request_translation,omitempty"`
}

// SpecialRequestAPI Special Request API
type SpecialRequestAPI struct {
	SpecialRequestCode *string `json:"special_request_code,omitempty" gorm:"type:varchar(36);not null;index:unique_special_request__special_request_code,unique,where:deleted_at is null"`
	SpecialRequestName *string `json:"special_request_name,omitempty" gorm:"type:varchar(256);not null;index:unique_special_request__special_request_name,unique,where:deleted_at is null"`
}
