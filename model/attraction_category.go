package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// AttractionCategory Attraction Category
type AttractionCategory struct {
	basic.Base
	basic.DataOwner
	AttractionCategoryAPI
	AttractionCategoryTranslation *AttractionCategoryTranslation `json:"attraction_category_translation,omitempty"` // Attraction Category Translation
	AttractionCategoryAsset       *AttractionCategoryAsset       `json:"attraction_category_asset,omitempty"`       // Attraction Category Asset
}

// AttractionCategoryAPI Attraction Category API
type AttractionCategoryAPI struct {
	AttractionCategoryName     *string                     `json:"attraction_category_name,omitempty" gorm:"type:varchar(256);not null;index:idx_attraction_category_name_deleted_at,unique,where:deleted_at is null" example:"Museums"` // Attraction Category Name
	ParentAttractionCategoryID *uuid.UUID                  `json:"parent_attraction_category_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`                                                                   // Parent Attraction Category ID
	IsDefault                  *bool                       `json:"is_default,omitempty"`                                                                                                                                                 // Is Default
	Description                *string                     `json:"description,omitempty" gorm:"type:text"`                                                                                                                               // Description
	AttractionCategoryAsset    *AttractionCategoryAssetAPI `json:"attraction_category_asset,omitempty" gorm:"-"`
}

// Seed Attraction Category
func (attractionCategory *AttractionCategory) Seed() *AttractionCategory {
	seed := AttractionCategory{
		AttractionCategoryAPI: AttractionCategoryAPI{
			AttractionCategoryName: lib.Strptr("Museum"),
		},
	}
	_ = lib.Merge(seed, &attractionCategory)
	return attractionCategory
}
