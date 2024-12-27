package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// QueueTypeTranslation Queue Type Translation
type QueueTypeTranslation struct {
	basic.Base
	basic.DataOwner
	QueueTypeTranslationAPI
	QueueTypeID *uuid.UUID `json:"queue_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:queue_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Queue Type id
}

// QueueTypeTranslationAPI Queue Type Translation API
type QueueTypeTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:queue_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	QueueTypeName *string `json:"queue_type_name,omitempty" gorm:"type:varchar(256)" example:"After Save Corporate "`                             // Queue Type Name
}
