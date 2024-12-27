package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// DerivedRateType Derived Rate Type
type DerivedRateType struct {
	basic.Base
	basic.DataOwner
	DerivedRateTypeAPI
	DerivedRateTypeTranslation *DerivedRateTypeTranslation `json:"derived_rate_type_translation,omitempty"` // Derived Rate Type Translation
}

// DerivedRateTypeAPI Derived Rate Type API
type DerivedRateTypeAPI struct {
	DerivedRateTypeCode *int    `json:"derived_rate_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null"`     // Derived Rate Type Code
	DerivedRateTypeName *string `json:"derived_rate_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Derived Rate Type Name
}
