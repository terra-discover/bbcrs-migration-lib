package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateHotel Model
type MarkupRateHotel struct {
	basic.Base
	basic.DataOwner
	MarkupRateHotelAPI
	MarkupRate *MarkupRate `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	Hotel      *Hotel      `json:"hotel" gorm:"foreignKey:HotelID;references:ID"`
}

// MarkupRateHotelAPI API
type MarkupRateHotelAPI struct {
	MarkupRateID *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	HotelID      *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded   *bool      `json:"is_included,omitempty"`
}
