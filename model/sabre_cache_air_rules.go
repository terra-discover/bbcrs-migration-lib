package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SabreCacheRules caching air rules
type SabreCacheAirRules struct {
	basic.Base
	basic.DataOwner
	SessionID                     *uuid.UUID       `json:"session_id" gorm:"type:varchar(36)" format:"uuid"` // session ID
	FlightID                      *uuid.UUID       `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Status                        *string          `json:"status" gorm:"type:varchar(60)"`
	DuplicateFareInfoText         *string          `json:"duplicate_fare_info_text" gorm:"type:text"`
	Amount                        *float64         `json:"amount" gorm:"type:decimal(22,2)"`
	CurrencyCode                  *string          `json:"currency_code" gorm:"type:varchar(3)"`
	Discontinue                   *string          `json:"discontinue" gorm:"type:varchar(60)"`
	Effective                     *string          `json:"effective" gorm:"type:varchar(60)"`
	FareClass                     *string          `json:"fare_class" gorm:"type:varchar(60)"`
	RoutingNumberOrMPM            *string          `json:"routing_number_or_mpm" gorm:"type:varchar(60)"`
	TariffDescriptionNumber       *string          `json:"tariff_description_number" gorm:"type:varchar(60)"`
	DisplayCode                   *string          `json:"display_code" gorm:"type:varchar(60)"`
	FareBasisCode                 *string          `json:"fare_basis_code" gorm:"type:varchar(60)"`
	FareVendor                    *string          `json:"fare_vendor" gorm:"type:varchar(60)"`
	FareText                      *string          `json:"fare_text" gorm:"type:text"`
	FareTypeCode                  *string          `json:"fare_type_code" gorm:"type:text"`
	FareTypeText                  *string          `json:"fare_type_text" gorm:"type:text"`
	FareType                      *string          `json:"fare_type" gorm:"type:text"`
	Footnotes                     *string          `json:"footnotes" gorm:"type:varchar(60)"`
	AirlineCode                   *string          `json:"airline_code" gorm:"type:varchar(60)"`
	OriginLocationCode            *string          `json:"origin_location_code" gorm:"type:varchar(60)"`
	DestinationLocationCode       *string          `json:"destination_location_code" gorm:"type:varchar(60)"`
	Rule                          *string          `json:"rule" gorm:"type:varchar(60)"`
	OriginTariffDescriptionNumber *string          `json:"origin_tariff_description_number" gorm:"type:varchar(60)"`
	TravelDate                    *string          `json:"travel_date" gorm:"type:varchar(60)"`
	AutoPrice                     *string          `json:"auto_price" gorm:"type:varchar(60)"`
	PassengerTypeCode             *string          `json:"passenger_type_code" gorm:"type:varchar(60)"`
	DBECode                       *string          `json:"dbe_code" gorm:"type:varchar(60)"`
	FareQuality                   *int64           `json:"fare_quality" gorm:"type:varchar(60)"`
	RouteCode                     *int64           `json:"route_code" gorm:"type:varchar(60)"`
	TariffFamily                  *string          `json:"tariff_family" gorm:"type:varchar(60)"`
	CreateDateTime                *strfmt.DateTime `json:"create_date_time" gorm:"type:varchar(60)"`
	ExpireDateTime                *string          `json:"expire_date_time" gorm:"type:varchar(60)"`
	RoutingInfoText               *string          `json:"routing_info_text" gorm:"type:text"`
	// Result                        *[]SabreCacheAirRulesStatus `json:"result,omitempty" gorm:"foreignKey:SabreCacheAirRulesID;references:ID"`
	MessageHeader *[]SabreCacheAirRulesHeader `json:"message_header,omitempty" gorm:"foreignKey:SabreCacheAirRulesID;references:ID"`
	Rules         *[]SabreCacheAirRulesRules  `json:"rules,omitempty" gorm:"foreignKey:SabreCacheAirRulesID;references:ID"`
}

type SabreCacheRules struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID      *uuid.UUID `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FareID        *uuid.UUID `json:"fare_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Type          *string    `json:"type,omitempty" gorm:"type:varchar(10)"`
	Applicability *string    `json:"applicability,omitempty" gorm:"type:varchar(10)"`
	RBD           *string    `json:"rbd,omitempty" gorm:"type:varchar(1)"`
	Amount        *int       `json:"Amount,omitempty"`
	Index         *int       `json:"index,omitempty"`
	Rule          *string    `json:"rule,omitempty" gorm:"type:text"`
}
