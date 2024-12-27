package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RecurrenceTypeTranslation Recurrence Type Translation
type RecurrenceTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RecurrenceTypeTranslationAPI
	RecurrenceTypeID *uuid.UUID `json:"recurrence_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:recurrence_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Recurrence Type id
}

// RecurrenceTypeTranslationAPI Recurrence Type Translation API
type RecurrenceTypeTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:recurrence_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RecurrenceTypeName *string `json:"recurrence_type_name,omitempty" gorm:"type:varchar(256)" example:"Daily"`                                             // Recurrence Type Name
}
