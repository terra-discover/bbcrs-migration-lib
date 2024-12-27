package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DerivedRateTypeTranslation Derived Rate Type Translation
type DerivedRateTypeTranslation struct {
	basic.Base
	basic.DataOwner
	DerivedRateTypeTranslationAPI
	DerivedRateTypeID *uuid.UUID `json:"derived_rate_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:derived_rate_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Derived Rate Type id
}

// DerivedRateTypeTranslationAPI Derived Rate Type Translation API
type DerivedRateTypeTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:derived_rate_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DerivedRateTypeName *string `json:"derived_rate_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Derived Rate Type Name
}
