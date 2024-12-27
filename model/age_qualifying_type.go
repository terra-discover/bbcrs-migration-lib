package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// AgeQualifyingType Age Qualifying Type
type AgeQualifyingType struct {
	basic.Base
	basic.DataOwner
	AgeQualifyingTypeAPI
	AgeQualifyingTypeTranslation *AgeQualifyingTypeTranslation `json:"age_qualifying_type_translation,omitempty"`
}

// AgeQualifyingTypeAPI Age Qualifying Type API
type AgeQualifyingTypeAPI struct {
	AgeQualifyingTypeCode *int64  `json:"age_qualifying_type_code,omitempty" gorm:"type:smallint;uniqueIndex:idx_age_qualifying_type_code_deleted_at,where:deleted_at is null;not null" example:"1"`                   // Age Qualifying Type Code
	AgeQualifyingTypeName *string `json:"age_qualifying_type_name,omitempty" gorm:"type:varchar(256);not null;index:unique_age_qualifying_type__age_qualifying_type_name,unique,where:deleted_at is null" example:"5"` // Age Qualifying Type Name
}
