package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateTravelPurpose Corporate Travel Purpose
type CorporateTravelPurpose struct {
	basic.Base
	basic.DataOwner
	CorporateTravelPurposeAPI
	Corporate     *Corporate     `json:"corporate,omitempty"`
	TravelPurpose *TravelPurpose `json:"travel_purpose,omitempty"`
}

// CorporateTravelPurposeAPI Corporate Travel Purpose API
type CorporateTravelPurposeAPI struct {
	TravelPurposeID *uuid.UUID `json:"travel_purpose_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CorporateID     *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IsDefault       *bool      `json:"is_default,omitempty"`
}
