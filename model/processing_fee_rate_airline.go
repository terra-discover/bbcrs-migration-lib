package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateAirline Model
type ProcessingFeeRateAirline struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateAirlineAPI
	ProcessingFeeRate *ProcessingFeeRate `json:"processing_fee_rate" gorm:"foreignKey:ProcessingFeeRateID;references:ID"`
	Airline           *Airline           `json:"airline" gorm:"foreignKey:AirlineID;references:ID"`
}

// ProcessingFeeRateAirlineAPI API
type ProcessingFeeRateAirlineAPI struct {
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	AirlineID           *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded          *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
