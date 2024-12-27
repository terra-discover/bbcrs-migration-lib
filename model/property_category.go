package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// PropertyCategory Property Category
type PropertyCategory struct {
	basic.Base
	basic.DataOwner
	PropertyCategoryAPI
	PropertyCategoryTranslation *PropertyCategoryTranslation `json:"property_category_translation,omitempty"`
}

// PropertyCategoryAPI Property Category API
type PropertyCategoryAPI struct {
	PropertyCategoryCode *int    `json:"property_category_code,omitempty" gorm:"type:smallint;uniqueIndex:idx_property_category_code_deleted_at,where:deleted_at is null;not null"`                 // Property Category Code
	PropertyCategoryName *string `json:"property_category_name,omitempty" gorm:"type:varchar(256);not null;index:unique_property_category__property_category_name,unique,where:deleted_at is null"` // Property Category Name
}
