package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BuiltInTypeTranslation Built In Type Translation
type BuiltInTypeTranslation struct {
	basic.Base
	basic.DataOwner
	BuiltInTypeTranslationAPI
	BuiltInTypeID *uuid.UUID `json:"built_in_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:built_in_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Built In Type id
}

// BuiltInTypeTranslationAPI Built In Type Translation API
type BuiltInTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:built_in_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	BuiltInTypeName *string `json:"built_in_type_name,omitempty" gorm:"type:varchar(256)" example:"Internal"`                                          // Built In Type Name
}
