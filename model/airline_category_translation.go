package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirlineCategoryTranslation Airline Category Translation
type AirlineCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	AirlineCategoryTranslationAPI
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:airline_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Airline Category id
}

// AirlineCategoryTranslationAPI Airline Category Translation API
type AirlineCategoryTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:airline_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AirlineCategoryName *string `json:"airline_category_name,omitempty" gorm:"type:varchar(64)" example:"Full Service"`                                       // Airline Category Name
}
