package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DimensionCategoryTranslation struct
type DimensionCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	DimensionCategoryName *string    `json:"dimension_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                                 // Dimension Category Name
	LanguageCode          *string    `json:"language_code,omitempty" gorm:"type:varchar(2);index:dimension_category_unique,unique,where:deleted_at is null" example:"en"`                                // 2 characters language code
	DimensionCategoryID   *uuid.UUID `json:"dimension_category_id,omitempty" gorm:"type:varchar(36);index:dimension_category_unique,unique,where:deleted_at is null" format:"uuid" swaggertype:"string"` // Dimension Category ID
}
