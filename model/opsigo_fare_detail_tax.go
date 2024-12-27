package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type OpsigoFareDetailTax struct {
	basic.Base
	basic.DataOwner
	OpsigoFareDetailPriceID *uuid.UUID `json:"opsigo_fare_detail_price_id"`
	TaxCode                 *string    `json:"tax_code,omitempty" gorm:"type:varchar(5)"`
	TaxName                 *string    `json:"tax_name,omitempty" gorm:"type:varchar(100)"`
	Amount                  *float64   `json:"amount,omitempty"`
	CountryCode             *string    `json:"country_code,omitempty" gorm:"type:varchar(5)"`
	CurrencyCode            *string    `json:"currency_code,omitempty" gorm:"type:varchar(36)"`
	CurrencyName            *string    `json:"currencey_name,omitempty" gorm:"type:varchar(50)"`
	Description             *string    `json:"description,omitempty" gorm:"type:text"`
}
