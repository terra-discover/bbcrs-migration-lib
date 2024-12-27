package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CabinTypeTranslation Cabin Type Translation
type CabinTypeTranslation struct {
	basic.Base
	basic.DataOwner
	CabinTypeTranslationAPI
	CabinTypeID *uuid.UUID `json:"cabin_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:cabin_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Cabin Type id
}

// CabinTypeTranslationAPI Cabin Type Translation API
type CabinTypeTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:cabin_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CabinTypeName *string `json:"cabin_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Cabin Type Name
}
