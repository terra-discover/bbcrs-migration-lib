package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateAirlineCategory Model
type MarkupRateAirlineCategory struct {
	basic.Base
	basic.DataOwner
	MarkupRateAirlineCategoryAPI
	MarkupRate      *MarkupRate      `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	AirlineCategory *AirlineCategory `json:"airline_category" gorm:"foreignKey:AirlineCategoryID;references:ID"`
}

// MarkupRateAirlineCategoryAPI API
type MarkupRateAirlineCategoryAPI struct {
	MarkupRateID      *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded        *bool      `json:"is_included,omitempty"`
}
