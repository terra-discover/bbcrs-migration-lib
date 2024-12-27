package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateAirline Model
type ServiceFeeRateAirline struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateAirlineAPI
	ServiceFeeRate *ServiceFeeRate `json:"service_fee_rate" gorm:"foreignKey:ServiceFeeRateID;references:ID"`
	Airline        *Airline        `json:"airline" gorm:"foreignKey:AirlineID;references:ID"`
}

// ServiceFeeRateAirlineAPI API
type ServiceFeeRateAirlineAPI struct {
	ServiceFeeRateID *uuid.UUID `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_airline__airline_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	AirlineID        *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_airline__airline_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	IsIncluded       *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
