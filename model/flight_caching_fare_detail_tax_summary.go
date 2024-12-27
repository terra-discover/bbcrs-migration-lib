package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingFareDetailTaxSummary struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_fare_detail_tax_summary__session_id;not null"`
	FlightCachingFareDetailTaxSummaryAPI
}

type FlightCachingFareDetailTaxSummaryAPI struct {
	FlightCachingFareDetailID *uuid.UUID `json:"flight_caching_fare_detail_id" gorm:"type:varchar(36)" format:"uuid" swaggertype:"string"`
	TaxCode                   *string    `json:"tax_code,omitempty" gorm:"type:varchar(5)"`
	TaxName                   *string    `json:"tax_name,omitempty" gorm:"type:varchar(100)"`
	Amount                    *float64   `json:"amount,omitempty"`
	CountryCode               *string    `json:"country_code,omitempty" gorm:"type:varchar(5)"`
	CurrencyCode              *string    `json:"currency_code,omitempty" gorm:"type:varchar(36)"`
	CurrencyName              *string    `json:"currencey_name,omitempty" gorm:"type:varchar(50)"`
	Description               *string    `json:"description,omitempty" gorm:"type:text"`
}
