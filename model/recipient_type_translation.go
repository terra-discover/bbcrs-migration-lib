package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RecipientTypeTranslation Recipient Type Translation
type RecipientTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RecipientTypeTranslationAPI
	RecipientTypeID *uuid.UUID `json:"recipient_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:recipient_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Recipient Type id
}

// RecipientTypeTranslationAPI Recipient Type Translation API
type RecipientTypeTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:recipient_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RecipientTypeName *string `json:"recipient_type_name,omitempty" gorm:"type:varchar(256)" example:"Agent/ Internal"`                                   // Recipient Type Name
}
