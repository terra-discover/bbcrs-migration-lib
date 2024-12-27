package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateHotelSupplier Model
type MarkupRateHotelSupplier struct {
	basic.Base
	basic.DataOwner
	MarkupRateHotelSupplierAPI
	MarkupRate    *MarkupRate    `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	HotelSupplier *HotelSupplier `json:"hotel_supplier" gorm:"foreignKey:HotelSupplierID;references:ID"`
}

// MarkupRateHotelSupplierAPI API
type MarkupRateHotelSupplierAPI struct {
	MarkupRateID    *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded      *bool      `json:"is_included,omitempty"`
}
