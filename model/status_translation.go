package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// StatusTranslation Status Translation
type StatusTranslation struct {
	basic.Base
	basic.DataOwner
	StatusTranslationAPI
	StatusID *uuid.UUID `json:"status_id,omitempty" gorm:"type:varchar(36);uniqueIndex:status_translation_unique;not null" swaggertype:"string" format:"uuid"` // Status id
}

// StatusTranslationAPI Status Translation API
type StatusTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:status_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	StatusName   *string `json:"status_name,omitempty" gorm:"type:varchar(256)" example:"Cancel/confirmed/requested "`                       // Status Name
}
