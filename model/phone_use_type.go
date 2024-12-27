package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// PhoneUseType Phone Use Type
type PhoneUseType struct {
	basic.Base
	basic.DataOwner
	PhoneUseTypeAPI
	PhoneUseTypeTranslation *PhoneUseTypeTranslation `json:"phone_use_type_translation,omitempty"` // Phone Use Type Translation
}

// PhoneUseTypeAPI Phone Use Type API
type PhoneUseTypeAPI struct {
	PhoneUseTypeCode *int    `json:"phone_use_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"`           // Phone Use Type Code
	PhoneUseTypeName *string `json:"phone_use_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Contact"` // Phone Use Type Name
}
