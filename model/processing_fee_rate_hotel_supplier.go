package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateHotelSupplier Model
type ProcessingFeeRateHotelSupplier struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateHotelSupplierAPI
	ProcessingFeeRate *ProcessingFeeRate `json:"processing_fee_rate" gorm:"foreignKey:ProcessingFeeRateID;references:ID"`
	HotelSupplier     *HotelSupplier     `json:"hotel_supplier" gorm:"foreignKey:HotelSupplierID;references:ID"`
}

// ProcessingFeeRateHotelSupplierAPI API
type ProcessingFeeRateHotelSupplierAPI struct {
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	HotelSupplierID     *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded          *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
