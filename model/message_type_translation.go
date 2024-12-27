package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MessageTypeTranslation Message Type Translation
type MessageTypeTranslation struct {
	basic.Base
	basic.DataOwner
	MessageTypeTranslationAPI
	MessageTypeID *uuid.UUID `json:"message_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:message_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Message Type id
}

// MessageTypeTranslationAPI Message Type Translation API
type MessageTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:message_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MessageTypeName *string `json:"message_type_name,omitempty" gorm:"type:varchar(256)" example:"Feedback"`                                          // Message Type Name
}
