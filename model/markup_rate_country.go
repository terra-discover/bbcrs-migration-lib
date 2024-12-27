package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateCountry Model
type MarkupRateCountry struct {
	basic.Base
	basic.DataOwner
	MarkupRateCountryAPI
	MarkupRate *MarkupRate `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	Country    *Country    `json:"country" gorm:"foreignKey:CountryID;references:ID"`
}

// MarkupRateCountryAPI API
type MarkupRateCountryAPI struct {
	MarkupRateID *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	CountryID    *uuid.UUID `json:"country_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded   *bool      `json:"is_included,omitempty"`
}
