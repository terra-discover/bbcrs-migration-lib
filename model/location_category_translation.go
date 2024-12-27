package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LocationCategoryTranslation Location Category Translation
type LocationCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	LocationCategoryTranslationAPI
	LocationCategoryID *uuid.UUID `json:"location_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:location_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Location Category id
}

// LocationCategoryTranslationAPI Location Category Translation API
type LocationCategoryTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:location_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	LocationCategoryName *string `json:"location_category_name,omitempty" gorm:"type:varchar(256)" example:"plaza semanggi"`                                    // Location Category Name
}
