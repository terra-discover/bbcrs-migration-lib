package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelBrand Hotel Brand
type HotelBrand struct {
	basic.Base
	basic.DataOwner
	HotelBrandAPI
	HotelBrandTranslation         *HotelBrandTranslation         `json:"hotel_brand_translation,omitempty"`
	CorporatePreferenceHotelBrand *CorporatePreferenceHotelBrand `json:"-"`
}

// HotelBrandAPI Hotel Brand API
type HotelBrandAPI struct {
	HotelBrandCode *string `json:"hotel_brand_code,omitempty" gorm:"type:varchar(36);not null;index:idx_hotel_brand_code_deleted_at,unique,where:deleted_at is null" example:"Hilton"`                    // Hotel Brand Code
	HotelBrandName *string `json:"hotel_brand_name,omitempty" gorm:"type:varchar(256);not null;index:idx_hotel_brand_code_deleted_at,unique,where:deleted_at is null" example:"Hilton Hotel and Resorts"` // Hotel Brand Name
}

// Seed HotelBrand
func (hotelBrand *HotelBrand) Seed() *HotelBrand {
	seed := HotelBrand{
		HotelBrandAPI: HotelBrandAPI{
			HotelBrandCode: lib.Strptr("Hilton"),
			HotelBrandName: lib.Strptr("Hilton Hotel and Resorts"),
		},
	}
	_ = lib.Merge(seed, &hotelBrand)
	return hotelBrand
}
