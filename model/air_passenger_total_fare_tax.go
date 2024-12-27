package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirPassengerTotalFareTax total fare tax
type AirPassengerTotalFareTax struct {
	basic.Base
	basic.DataOwner
	AirPassengerTotalFareTaxAPI
	AirPassengerTotalFare *AirPassengerTotalFare `json:"air_passenger_total_fare,omitempty" gorm:"foreignKey:AirPassengerTotalFareID;references:ID"`
}

type AirPassengerTotalFareTaxAPI struct {
	AirPassengerTotalFareID *uuid.UUID `json:"air_passenger_total_fare_id,omitempty" gorm:"type:varchar(36);not null"`
	TaxCode                 *string    `json:"tax_code,omitempty" gorm:"type:varchar(50)"`
	TaxName                 *string    `json:"tax_name,omitempty" gorm:"type:varchar(256)"`
	CurrencyID              *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)"`
	Amount                  *float64   `json:"amount,omitempty"`
	PublishedCurrencyID     *string    `json:"published_currency_id,omitempty" gorm:"type:varchar(36)"`
	PublishedAmount         *float64   `json:"published_amount,omitempty"`
	Station                 *string    `json:"station,omitempty" gorm:"type:varchar(8)"`
	CountryID               *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36)"`
	Description             *string    `json:"description,omitempty" gorm:"type:text"`
}
