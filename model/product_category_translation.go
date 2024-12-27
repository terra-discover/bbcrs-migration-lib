package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProductCategoryTranslation Product Category Translation
type ProductCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	ProductCategoryTranslationAPI
	ProductCategoryID *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:product_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Product category id
}

// ProductCategoryTranslationAPI Product Category Translation API
type ProductCategoryTranslationAPI struct {
	LanguageID          *uuid.UUID `json:"language_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProductCategoryName *string    `json:"product_category_name,omitempty" gorm:"type:varchar(256)"` // Product Category Name
	Description         *string    `json:"description,omitempty" gorm:"type:varchar(4000)"`          // Description
}
