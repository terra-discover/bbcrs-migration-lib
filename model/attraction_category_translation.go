package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AttractionCategoryTranslation Attraction Category Translation
type AttractionCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	AttractionCategoryTranslationAPI
	AttractionCategoryID *uuid.UUID `json:"attraction_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:attraction_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Attraction Category id
}

// AttractionCategoryTranslationAPI Attraction Category Translation API
type AttractionCategoryTranslationAPI struct {
	LanguageCode           *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:attraction_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AttractionCategoryName *string `json:"attraction_category_name,omitempty" gorm:"type:varchar(256)" example:"Museums"`                                           // Attraction Category Name
	Description            *string `json:"description,omitempty" gorm:"type:text"`                                                                                  // Description
}
