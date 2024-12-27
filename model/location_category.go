package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// LocationCategory Location Category
type LocationCategory struct {
	basic.Base
	basic.DataOwner
	LocationCategoryAPI
	LocationCategoryTranslation *LocationCategoryTranslation `json:"location_category_translation,omitempty"`
}

// LocationCategoryAPI Location Category API
type LocationCategoryAPI struct {
	LocationCategoryCode *int    `json:"location_category_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"1"`                  // Location Category Code
	LocationCategoryName *string `json:"location_category_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"plaza semanggi"` // Location Category Name
}
