package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type OpsigoFareDetailPrice struct {
	basic.Base
	basic.DataOwner
	SessionID           *uuid.UUID             `json:"session_id" gorm:"type:varchar(36)"`
	PassengerType       *string                `json:"passenger_type" gorm:"type:varchar(30)"`
	FareID              *uuid.UUID             `json:"fare_id" gorm:"type:varchar(36)"`
	ClassID             *string                `json:"class_id" gorm:"type:varchar(1000)"`
	FlightID            *string                `json:"flight_id" gorm:"type:varchar(1000)"`
	Price               *float64               `json:"price"`
	TaxPrice            *float64               `json:"tax_price"`
	OpsigoFareDetailTax *[]OpsigoFareDetailTax `json:"opsigo_fare_detail_tax" gorm:"foreignKey:OpsigoFareDetailPriceID;references:ID"`
}
