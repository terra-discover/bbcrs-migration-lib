package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CurrencyTranslation Currency Translation
type CurrencyTranslation struct {
	basic.Base
	basic.DataOwner
	CurrencyTranslationAPI
	CurrencyID *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36);uniqueIndex:currency_translation_unique;not null" swaggertype:"string" format:"uuid"` // Currency id
}

// CurrencyTranslationAPI Currency Translation API
type CurrencyTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:currency_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CurrencyName   *string `json:"currency_name,omitempty" gorm:"type:varchar(64)" example:"Rupiah Indonesia"`                                   // Currency Name
	CurrencySymbol *string `json:"currency_symbol,omitempty" gorm:"type:varchar(64)" example:"Rp"`                                               // Currency Symbol
	MinorUnitName  *string `json:"minor_unit_name,omitempty" gorm:"type:varchar(64)" example:"Rupiah Indonesia"`                                 // Minor Unit Name
}
