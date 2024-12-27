package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// GuaranteeTypeTranslation Guarantee Type Translation
type GuaranteeTypeTranslation struct {
	basic.Base
	basic.DataOwner
	GuaranteeTypeTranslationAPI
	GuaranteeTypeID *uuid.UUID `json:"guarantee_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:guarantee_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Guarantee Type id
}

// GuaranteeTypeTranslationAPI Guarantee Type Translation API
type GuaranteeTypeTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:guarantee_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	GuaranteeTypeName *string `json:"guarantee_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Guarantee Type Name
}
