package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FeeTaxTypeTranslation Fee Tax Type Translation
type FeeTaxTypeTranslation struct {
	basic.Base
	basic.DataOwner
	FeeTaxTypeTranslationAPI
	FeeTaxTypeID *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:fee_tax_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Flight Type id
}

// FeeTaxTypeTranslationAPI Fee Tax Type Translation API
type FeeTaxTypeTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:fee_tax_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FeeTaxTypeName *string `json:"fee_tax_type_name,omitempty" gorm:"type:varchar(256)" example:"Forces"`                                            // Fee Tax Type Name
	Description    *string `json:"description,omitempty" gorm:"type:text" example:"Forces"`                                                          // Description
}
