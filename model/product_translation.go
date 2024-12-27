package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProductTranslation Product Translation
type ProductTranslation struct {
	basic.Base
	basic.DataOwner
	ProductTranslationAPI
	ProductID *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);uniqueIndex:product_translation_unique;not null" swaggertype:"string" format:"uuid"` // Product id
}

// ProductTranslationAPI Product Translation API
type ProductTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:product_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ProductName  *string `json:"product_name,omitempty" gorm:"type:varchar(256)"`                                                             // Product Name
	ProductUnit  *string `json:"product_unit,omitempty" gorm:"type:varchar(32)"`                                                              // Product Unit
	Description  *string `json:"description,omitempty" gorm:"type:varchar(4000)" example:"description product"`                               // Description
}
