package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FieldTypeTranslation Field Type Translation
type FieldTypeTranslation struct {
	basic.Base
	basic.DataOwner
	FieldTypeTranslationAPI
	FieldTypeID *uuid.UUID `json:"field_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:field_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Field Type id
}

// FieldTypeTranslationAPI Field Type Translation API
type FieldTypeTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:field_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FieldTypeName *string `json:"field_type_name,omitempty" gorm:"type:varchar(256)" example:"radio button"`                                      // Field Type Name
}
