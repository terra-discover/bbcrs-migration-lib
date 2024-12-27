package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateCity Model
type MarkupRateCity struct {
	basic.Base
	basic.DataOwner
	MarkupRateCityAPI
	MarkupRate *MarkupRate `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	City       *City       `json:"city" gorm:"foreignKey:CityID;references:ID"`
}

// MarkupRateCityAPI API
type MarkupRateCityAPI struct {
	MarkupRateID *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	CityID       *uuid.UUID `json:"city_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded   *bool      `json:"is_included,omitempty"`
}
