package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelSupplier Hotel Supplier
type HotelSupplier struct {
	basic.Base
	basic.DataOwner
	HotelSupplierAPI
	SupplierType             *SupplierType             `json:"supplier_type,omitempty"`
	HotelSupplierTranslation *HotelSupplierTranslation `json:"hotel_supplier_translation,omitempty"`
}

// HotelSupplierAPI Hotel Supplier API
type HotelSupplierAPI struct {
	HotelSupplierCode *string    `json:"hotel_supplier_code,omitempty" gorm:"type:varchar(36);not null;index:idx_hotel_supplier_code_deleted_at,unique,where:deleted_at is null" example:"BDC"`          // Hotel Supplier Code
	HotelSupplierName *string    `json:"hotel_supplier_name,omitempty" gorm:"type:varchar(256);not null;index:idx_hotel_supplier_name_deleted_at,unique,where:deleted_at is null" example:"Booking.com"` // Hotel Supplier Name
	SupplierTypeID    *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                                                                               // Supplier Type ID
}

// Seed HotelSupplier
func (hotelSupplier *HotelSupplier) Seed() *HotelSupplier {
	seed := HotelSupplier{
		HotelSupplierAPI: HotelSupplierAPI{
			HotelSupplierCode: lib.Strptr("BDC"),
			HotelSupplierName: lib.Strptr("Booking.com"),
			SupplierTypeID:    lib.UUIDPtr(uuid.New()),
		},
	}
	_ = lib.Merge(seed, &hotelSupplier)
	return hotelSupplier
}
