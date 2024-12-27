package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirItineraryTotalFareFee struct {
	basic.Base
	basic.DataOwner
	AirItineraryTotalFareFeeAPI
}

type AirItineraryTotalFareFeeAPI struct {
	AirItineraryTotalFareID *uuid.UUID `json:"air_itinerary_total_fare_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	FeeCode                 *string    `json:"fee_code,omitempty" gorm:"type:varchar(36)"`
	FeeName                 *string    `json:"fee_name,omitempty" gorm:"type:varchar(256)"`
	CurrencyID              *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" type:"uuid"`
	Amount                  *float64   `json:"amount,omitempty"`
	Description             *string    `json:"description,omitempty" gorm:"type:text"`
}
