package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCurrencyTranslation Integration Partner Currency Translation
type IntegrationPartnerCurrencyTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCurrencyTranslationAPI
	IntegrationPartnerCurrencyID *uuid.UUID `json:"integration_partner_currency_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_currency_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Currency id
}

// IntegrationPartnerCurrencyTranslationAPI Integration Partner Currency Translation API
type IntegrationPartnerCurrencyTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_currency_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CurrencyName   *string `json:"currency_name,omitempty" gorm:"type:varchar(64)"`                                                                                  // Currency Name
	CurrencySymbol *string `json:"currency_symbol,omitempty" gorm:"type:varchar(64)"`                                                                                // Currency Symbol
}
