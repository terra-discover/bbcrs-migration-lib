package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// FieldType Field Type
type FieldType struct {
	basic.Base
	basic.DataOwner
	FieldTypeAPI
	FieldTypeTranslation *FieldTypeTranslation `json:"field_type_translation,omitempty"`
}

// FieldTypeAPI Field Type API
type FieldTypeAPI struct {
	FieldTypeCode *string `json:"field_type_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" example:"rbtn"`          // Field Type Code
	FieldTypeName *string `json:"field_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"radio button"` // Field Type Name
}
