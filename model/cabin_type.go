package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// CabinType Cabin Type
type CabinType struct {
	basic.Base
	basic.DataOwner
	CabinTypeAPI
	CabinTypeTranslation *CabinTypeTranslation `json:"cabin_type_translation,omitempty"`
}

// CabinTypeAPI Cabin Type API
type CabinTypeAPI struct {
	CabinTypeCode *string `json:"cabin_type_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null" validate:"required,gt=0,lte=36"`   // Cabin Type Code
	CabinTypeName *string `json:"cabin_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" validate:"required,gt=0,lte=256"` // Cabin Type Name
	IsDefault     *bool   `json:"is_default,omitempty"`                                                                                                                // Is Default
}
