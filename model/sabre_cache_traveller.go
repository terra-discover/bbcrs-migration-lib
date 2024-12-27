package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreCacheTravellers struct {
	basic.Base
	basic.DataOwner
	SessionID                *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)"`
	Index                    *int       `json:"index,omitempty" gorm:"type:int"`
	Type                     *string    `json:"type,omitempty" gorm:"type:varchar(60)"`
	GivenName                *string    `json:"given_name,omitempty" gorm:"type:varchar(60)"`
	Surname                  *string    `json:"surname,omitempty" gorm:"type:varchar(60)"`
	MiddleName               *string    `json:"middle_name,omitempty" gorm:"type:varchar(60)"`
	Gender                   *string    `json:"gender,omitempty" gorm:"type:varchar(10)"`
	IsInfant                 *int       `json:"is_infant,omitempty"`
	ParentNumber             *int       `json:"parent_number,omitempty"`
	PricePax                 *float64   `json:"price_pax,omitempty"`
	Currency                 *string    `json:"currency,omitempty"`
	PriceSeat                *float64   `json:"price_seat,omitempty"`
	PriceAddons              *float64   `json:"price_addons,omitempty"`
	TotalPrice               *float64   `json:"total_price,omitempty"`
	TicketDocumentNumber     *string    `json:"ticket_document_number,omitempty" gorm:"type:varchar(64)" swaggertype:"string"`
	BaseFareCurrencyID       *uuid.UUID `json:"base_fare_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	BaseFareAmount           *float64   `json:"base_fare_amount,omitempty"`
	EquivalentFareCurrencyID *uuid.UUID `json:"equivalent_fare_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	EquivalentFareAmount     *float64   `json:"equivalent_fare_amount,omitempty"`
	TotalTaxCurrencyID       *uuid.UUID `json:"total_tax_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	TotalTaxAmount           *float64   `json:"total_tax_amount,omitempty"`
	TotalFeeCurrencyID       *uuid.UUID `json:"total_fee_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	TotalFeeAmount           *float64   `json:"total_fee_amount,omitempty"`
	TotalFareCurrencyID      *uuid.UUID `json:"total_fare_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	TotalFareAmount          *float64   `json:"total_fare_amount,omitempty"`
	// AddOns       SabreSumAddons `json:"addons,omitempty" gorm:"foreignKey:TravellerID;references:ID"`
	// Seat         SabreSumSeat   `json:"seat,omitempty" gorm:"foreignKey:TravellerID;references:ID"`
}

type SabreSumAddons struct {
	TraverllerID *uuid.UUID `json:"travellerID,omitempty" gorm:"type:varchar(36)"`
	Type         *string    `json:"type,omitempty" gorm:"type:varchar(60)"`
	Name         *string    `json:"name,omitempty" gorm:"type:varchar(256)"`
	Info         *string    `json:"info,omitempty" gorm:"type:varchar(256)"`
	DepAirport   *string    `json:"dep_airport,omitempty" gorm:"type:varchar(3)"`
	ArrAirport   *string    `json:"arr_airport,omitempty" gorm:"type:varchar(3)"`
	Price        *float64   `json:"price,omitempty"`
	RficCode     *string    `json:"rfic_code,omitempty" gorm:"type:varchar(1)"`
	RficSubcode  *string    `json:"rfic_subcode,omitempty" gorm:"type:varchar(5)"`
}

type SabreSumSeat struct {
	TraverllerID *uuid.UUID `json:"travellerID,omitempty" gorm:"type:varchar(36)"`
	Type         *string    `json:"type,omitempty" gorm:"type:varchar(60)"`
	Name         *string    `json:"name,omitempty" gorm:"type:varchar(256)"`
	Info         *string    `json:"info,omitempty" gorm:"type:varchar(256)"`
	DepAirport   *string    `json:"dep_airport,omitempty" gorm:"type:varchar(3)"`
	ArrAirport   *string    `json:"arr_airport,omitempty" gorm:"type:varchar(3)"`
	Price        *float64   `json:"price,omitempty"`
	RficCode     *string    `json:"rfic_code,omitempty" gorm:"type:varchar(1)"`
	RficSubcode  *string    `json:"rfic_subcode,omitempty" gorm:"type:varchar(5)"`
}
