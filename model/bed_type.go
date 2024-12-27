package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// BedType Bed Type
type BedType struct {
	basic.Base
	basic.DataOwner
	BedTypeAPI
	BedTypeTranslation *BedTypeTranslation `json:"bed_type_translation,omitempty"`
}

// BedTypeAPI Bed Type API
type BedTypeAPI struct {
	BedTypeCode *int    `json:"bed_type_code,omitempty" gorm:"type:smallint;uniqueIndex:idx_bed_type_code_deleted_at,where:deleted_at is null;not null" example:"1"` // Bed Type Code
	BedTypeName *string `json:"bed_type_name,omitempty" gorm:"type:varchar(256);not null" example:"queen bed"`                                                       // Bed Type Name
}
