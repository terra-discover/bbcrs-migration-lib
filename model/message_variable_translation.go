package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MessageVariableTranslation Message Variable Translation
type MessageVariableTranslation struct {
	basic.Base
	basic.DataOwner
	MessageVariableTranslationAPI
	MessageVariableID *uuid.UUID `json:"message_variable_id,omitempty" gorm:"type:varchar(36);uniqueIndex:message_variable_translation_unique;not null" swaggertype:"string" format:"uuid"` // Message Variable id
}

// MessageVariableTranslationAPI Message Variable Translation API
type MessageVariableTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:message_variable_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MessageVariableName *string `json:"message_variable_name,omitempty" gorm:"type:varchar(256)" example:"Departure Date"`                                    // Message Variable Name
	DummyValue          *string `json:"dummy_value,omitempty" gorm:"type:text"`                                                                               // Dummy Value
	Description         *string `json:"description,omitempty" gorm:"type:text"`                                                                               // Description
}
