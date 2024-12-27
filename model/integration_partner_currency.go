package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCurrency Integration Partner Currency
type IntegrationPartnerCurrency struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCurrencyAPI
	IntegrationPartnerCurrencyTranslation *IntegrationPartnerCurrencyTranslation `json:"integration_partner_currency_translation,omitempty"` // Integration Partner Currency Translation
	Currency                              *Currency                              `json:"currency,omitempty"`                                 // Currency
}

// IntegrationPartnerCurrencyAPI Integration Partner Currency API
type IntegrationPartnerCurrencyAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"index:idx_idx_integration_partner_currency__currency_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	CurrencyID           *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" gorm:"type:varchar(36)" format:"uuid"`                                                                                              // Currency ID
	CurrencyCode         *string    `json:"currency_code,omitempty" gorm:"type:varchar(3);index:idx_idx_integration_partner_currency__currency_code,unique,where:deleted_at is null;not null"`                             // Currency Code
	CurrencyName         *string    `json:"currency_name,omitempty" gorm:"type:varchar(64)"`                                                                                                                               // Currency Name
	CurrencySymbol       *string    `json:"currency_symbol,omitempty" gorm:"type:varchar(64)"`                                                                                                                             // Currency Symbol
}
