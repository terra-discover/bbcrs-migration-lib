package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DocumentTypeTranslation Document Type Translation
type DocumentTypeTranslation struct {
	basic.Base
	basic.DataOwner
	DocumentTypeTranslationAPI
	DocumentTypeID *uuid.UUID `json:"document_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:document_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Document Type id
}

// DocumentTypeTranslationAPI Document Type Translation API
type DocumentTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:document_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DocumentTypeName *string `json:"document_type_name,omitempty" gorm:"type:varchar(256)" example:"pdf"`                                               // Document Type Name
}
