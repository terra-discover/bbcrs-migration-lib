package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteCurrency SiteCurrency
type SiteCurrency struct {
	basic.Base
	basic.DataOwner
	SiteCurrencyAPI
	Site     *Site     `json:"site,omitempty" gorm:"foreignKey:SiteID;references:ID"`
	Currency *Currency `json:"currency,omitempty" gorm:"foreignKey:CurrencyID;references:ID"`
}

// SiteCurrencyAPI SiteCurrency API
type SiteCurrencyAPI struct {
	SiteID     *uuid.UUID `json:"site_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	CurrencyID *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsPrimary  *bool      `json:"is_primary,omitempty"`
}
