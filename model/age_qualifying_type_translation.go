package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgeQualifyingTypeTranslation Age Qualifying Type Translation
type AgeQualifyingTypeTranslation struct {
	basic.Base
	basic.DataOwner
	AgeQualifyingTypeTranslationAPI
	AgeQualifyingTypeID *uuid.UUID `json:"age_qualifying_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:age_qualifying_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Age Qualifying Type id
}

// AgeQualifyingTypeTranslationAPI Age Qualifying Type Translation API
type AgeQualifyingTypeTranslationAPI struct {
	LanguageCode          *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:age_qualifying_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AgeQualifyingTypeName *string `json:"age_qualifying_type_name,omitempty" gorm:"type:varchar(256)" example:"5"`                                                 // Age Qualifying Type Name
}
