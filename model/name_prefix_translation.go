package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// NamePrefixTranslation Name Prefix Translation
type NamePrefixTranslation struct {
	basic.Base
	basic.DataOwner
	NamePrefixTranslationAPI
	NamePrefixID *uuid.UUID `json:"name_prefix_id,omitempty" gorm:"type:varchar(36);uniqueIndex:name_prefix_translation_unique;not null" swaggertype:"string" format:"uuid"` // Name Prefix id
}

// NamePrefixTranslationAPI Name Prefix Translation API
type NamePrefixTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:name_prefix_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	NamePrefixName *string `json:"name_prefix_name,omitempty" gorm:"type:varchar(256)"`                                                             // Name Prefix Name
}
