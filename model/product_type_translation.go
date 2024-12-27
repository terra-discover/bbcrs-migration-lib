package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProductTypeTranslation Product Type Translation
type ProductTypeTranslation struct {
	basic.Base
	basic.DataOwner
	ProductTypeTranslationAPI
	ProductTypeID *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:product_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Product Type id
}

// ProductTypeTranslationAPI Product Type Translation API
type ProductTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:product_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ProductTypeName *string `json:"product_type_name,omitempty" gorm:"type:varchar(256)" example:"Flight"`                                            // Product Type Name
}
