package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// RatingType Rating Type
type RatingType struct {
	basic.Base
	basic.DataOwner
	RatingTypeAPI
	RatingTypeTranslation *RatingTypeTranslation `json:"rating_type_translation,omitempty"`
}

// RatingTypeAPI Rating Type API
type RatingTypeAPI struct {
	RatingTypeCode *int    `json:"rating_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null"`     // Rating Type Code
	Scale          *int    `json:"scale,omitempty" gorm:"type:smallint"`                                                                // Scale
	RatingTypeName *string `json:"rating_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Rating Type Name
	Provider       *string `json:"provider,omitempty" gorm:"type:varchar(256)"`                                                         // Provider
	RatingSymbol   *string `json:"rating_symbol,omitempty" gorm:"type:varchar(256)"`                                                    // Rating Symbol
}
