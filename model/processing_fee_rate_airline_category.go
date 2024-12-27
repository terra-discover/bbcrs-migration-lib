package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateAirlineCategory Model
type ProcessingFeeRateAirlineCategory struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateAirlineCategoryAPI
	ProcessingFeeRate *ProcessingFeeRate `json:"processing_fee_rate" gorm:"foreignKey:ProcessingFeeRateID;references:ID"`
	AirlineCategory   *AirlineCategory   `json:"airline_category" gorm:"foreignKey:AirlineCategoryID;references:ID"`
}

// ProcessingFeeRateAirlineCategoryAPI API
type ProcessingFeeRateAirlineCategoryAPI struct {
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	AirlineCategoryID   *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded          *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
