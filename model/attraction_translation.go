package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AttractionTranslation Attraction Translation
type AttractionTranslation struct {
	basic.Base
	basic.DataOwner
	AttractionTranslationAPI
	AttractionID *uuid.UUID `json:"attraction_id,omitempty" gorm:"type:varchar(36);uniqueIndex:attraction_translation_unique;not null" swaggertype:"string" format:"uuid"` // Attraction id
}

// AttractionTranslationAPI Attraction Translation API
type AttractionTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:attraction_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AttractionName *string `json:"attraction_name,omitempty" gorm:"type:varchar(256)"`                                                             // Attraction Name
	Description    *string `json:"description,omitempty" gorm:"type:text"`                                                                         // Description
}
