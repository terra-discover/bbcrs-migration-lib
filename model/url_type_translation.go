package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// URLTypeTranslation URL Type Translation
type URLTypeTranslation struct {
	basic.Base
	basic.DataOwner
	URLTypeTranslationAPI
	URLTypeID *uuid.UUID `json:"url_type_id,omitempty" gorm:"type:varchar(36);not null;uniqueIndex:url_type_translation_unique" swaggertype:"string" format:"uuid"` // URL Type id
}

// URLTypeTranslationAPI URL Type Translation API
type URLTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);not null;uniqueIndex:url_type_translation_unique" example:"en"` // Language code example: en, id, cn etc...
	URLTypeName  *string `json:"url_type_name,omitempty" gorm:"type:varchar(256)" example:"Website"`                                           // URL Type Name
}
