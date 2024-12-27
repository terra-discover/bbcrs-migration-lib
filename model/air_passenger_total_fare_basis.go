package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirPassengerTotalFareBasis struct {
	basic.Base
	basic.DataOwner
	AirPassengerTotalFareBasisAPI
	AirPassengerTotalFare *AirPassengerTotalFare `json:"air_passenger_total_fare,omitempty" gorm:"foreignKey:AirPassengerTotalFareID;references:ID"`
	PassengerType         *PassengerType         `json:"passenger_type,omitempty" gorm:"foreignKey:PassengerTypeID;references:ID"`
}

type AirPassengerTotalFareBasisAPI struct {
	AirPassengerTotalFareID *uuid.UUID       `json:"air_passenger_total_fare_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	FlightSegmentID         *uuid.UUID       `json:"flight_segment_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FareBasisCode           *string          `json:"fare_basis_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	NotValidBefore          *strfmt.DateTime `json:"not_valid_before,omitempty" gorm:"type:timestamptz" format:"date-time"`
	NotValidAfter           *strfmt.DateTime `json:"not_valid_after,omitempty" gorm:"type:timestamptz" format:"date-time"`
	FareCurrencyID          *string          `json:"fare_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	FareAmount              *float64         `json:"fare_amount,omitempty"`
	PublishedFareCurrencyID *string          `json:"published_fare_currency_id,omitempty" gorm:"type:varchar(36)"`
	PublishedFareAmount     *float64         `json:"published_fare_amount,omitempty"`
	PassengerTypeID         *uuid.UUID       `json:"passenger_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	VendorCode              *string          `json:"vendor_code,omitempty" gorm:"type:varchar(16)"`
	FareType                *string          `json:"fare_type,omitempty" gorm:"type:varchar(16)"`
	FareTariff              *string          `json:"fare_tariff,omitempty" gorm:"type:varchar(16)"`
	FareCode                *string          `json:"fare_code,omitempty" gorm:"type:varchar(16)"`
	ClassRef                *string          `json:"class_ref,omitempty" gorm:"type:varchar(36)"`
	ClassCode               *string          `json:"class_code,omitempty" gorm:"type:varchar(36)"`
	ClassName               *string          `json:"class_name,omitempty" gorm:"type:varchar(36)"`
	IsPrivateFare           *bool            `json:"is_private_fare,omitempty"`
	TourCode                *string          `json:"tour_code,omitempty" gorm:"type:varchar(36)"`
}
