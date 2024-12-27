package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAmenityCategory Hotel Amenity Category
type HotelAmenityCategory struct {
	basic.Base
	basic.DataOwner
	HotelAmenityCategoryAPI
	HotelAmenityCategoryTranslation *HotelAmenityCategoryTranslation `json:"hotel_amenity_category_translation,omitempty"` // Hotel Amenity Category Translation
	HotelAmenityCategoryAsset       *HotelAmenityCategoryAsset       `json:"hotel_amenity_category_asset,omitempty"`       // Hotel Amenity Category Asset
}

// HotelAmenityCategoryAPI Hotel Amenity Category API
type HotelAmenityCategoryAPI struct {
	ParentHotelAmenityCategoryID *uuid.UUID                    `json:"parent_hotel_amenity_category_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // Hotel Amenity Category id
	HotelAmenityCategoryName     *string                       `json:"hotel_amenity_category_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Hotel Amenity Category Name
	IsDefault                    *bool                         `json:"is_default,omitempty"`                                                                                           // Is Default
	Description                  *string                       `json:"description,omitempty" gorm:"type:text"`                                                                         // Description
	HotelAmenityCategoryAsset    *HotelAmenityCategoryAssetAPI `json:"hotel_amenity_category_asset,omitempty" gorm:"-"`
}

// Seed HotelAmenityCategory Category
func (hotelAmenityCategory *HotelAmenityCategory) Seed() *HotelAmenityCategory {
	name := "Soap"
	description := "Soap Soft"
	isDefault := true
	parentUUID, _ := uuid.NewUUID()
	seed := HotelAmenityCategory{
		HotelAmenityCategoryAPI: HotelAmenityCategoryAPI{
			HotelAmenityCategoryName:     &name,
			Description:                  &description,
			IsDefault:                    &isDefault,
			ParentHotelAmenityCategoryID: &parentUUID,
		},
	}
	_ = lib.Merge(seed, &hotelAmenityCategory)
	return hotelAmenityCategory
}
