package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateHotelSupplier Model
type ServiceFeeRateHotelSupplier struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateHotelSupplierAPI
	ServiceFeeRate *ServiceFeeRate `json:"service_fee_rate" gorm:"foreignKey:ServiceFeeRateID;references:ID"`
	HotelSupplier  *HotelSupplier  `json:"hotel_supplier" gorm:"foreignKey:HotelSupplierID;references:ID"`
}

// ServiceFeeRateHotelSupplierAPI API
type ServiceFeeRateHotelSupplierAPI struct {
	ServiceFeeRateID *uuid.UUID `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_hotel_supplier__hotel_supplier_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	HotelSupplierID  *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_hotel_supplier__hotel_supplier_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	IsIncluded       *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
