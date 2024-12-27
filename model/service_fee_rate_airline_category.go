package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateAirlineCategory Model
type ServiceFeeRateAirlineCategory struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateAirlineCategoryAPI
	ServiceFeeRate  *ServiceFeeRate  `json:"service_fee_rate" gorm:"foreignKey:ServiceFeeRateID;references:ID"`
	AirlineCategory *AirlineCategory `json:"airline_category" gorm:"foreignKey:AirlineCategoryID;references:ID"`
}

// ServiceFeeRateAirlineCategoryAPI API
type ServiceFeeRateAirlineCategoryAPI struct {
	ServiceFeeRateID  *uuid.UUID `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_airline_category__airline_category_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_airline_category__airline_category_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	IsIncluded        *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
