package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageTypeTranslation Page Type Translation
type PageTypeTranslation struct {
	basic.Base
	basic.DataOwner
	PageTypeTranslationAPI
	PageTypeID *uuid.UUID `json:"page_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:page_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Page Type id
}

// PageTypeTranslationAPI Page Type Translation API
type PageTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:page_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PageTypeName *string `json:"page_type_name,omitempty" gorm:"type:varchar(256)" example:"Feedback"`                                          // Page Type Name
}
