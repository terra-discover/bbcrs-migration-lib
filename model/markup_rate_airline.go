package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateAirline Model
type MarkupRateAirline struct {
	basic.Base
	basic.DataOwner
	MarkupRateAirlineAPI
	MarkupRate *MarkupRate `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	Airline    *Airline    `json:"airline" gorm:"foreignKey:AirlineID;references:ID"`
}

// MarkupRateAirlineAPI API
type MarkupRateAirlineAPI struct {
	MarkupRateID *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	AirlineID    *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded   *bool      `json:"is_included,omitempty"`
}
