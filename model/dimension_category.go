package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// DimensionCategory struct
type DimensionCategory struct {
	basic.Base
	basic.DataOwner
	DimensionCategoryCode        *int64                        `json:"dimension_category_code,omitempty" gorm:"type:smallint;uniqueIndex:,where:deleted_at is null;not null"` // Dimension Category Code
	DimensionCategoryName        *string                       `json:"dimension_category_name,omitempty" gorm:"type:varchar(256);not null"`                                   // Dimension Category Name
	DimensionCategoryTranslation *DimensionCategoryTranslation `json:"dimension_category_translation,omitempty"`                                                              // Dimension Category Translation
}
