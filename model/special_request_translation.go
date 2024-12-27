package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SpecialRequestTranslation Special Request Translation
type SpecialRequestTranslation struct {
	basic.Base
	basic.DataOwner
	SpecialRequestTranslationAPI
	SpecialRequestID *uuid.UUID `json:"special_request_id,omitempty" gorm:"type:varchar(36);uniqueIndex:special_request_translation_unique;not null" swaggertype:"string" format:"uuid"` // Special Request id
}

// SpecialRequestTranslationAPI Special Request Translation API
type SpecialRequestTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:special_request_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SpecialRequestName *string `json:"special_request_name,omitempty" gorm:"type:varchar(256)"`                                                             // Special Request Name
}
