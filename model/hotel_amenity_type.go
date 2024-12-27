package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// HotelAmenityType Hotel Amenity Type
type HotelAmenityType struct {
	basic.Base
	basic.DataOwner
	HotelAmenityTypeAPI
	HotelAmenityTypeTranslation          *HotelAmenityTypeTranslation           `json:"hotel_amenity_type_translation,omitempty"`            // Hotel Amenity Type Translation
	HotelAmenityCategoryHotelAmenityType []HotelAmenityCategoryHotelAmenityType `json:"hotel_amenity_category_hotel_amenity_type,omitempty"` // Hotel Amenity Category Hotel Amenity Type
	HotelAmenityTypeAsset                *HotelAmenityTypeAsset                 `json:"hotel_amenity_type_asset,omitempty"`                  // Hotel Amenity Type Asset
}

// HotelAmenityTypeList Read Only Model
type HotelAmenityTypeList struct {
	HotelAmenityType
	HotelAmenityCategoryNames *string `json:"hotel_amenity_category_names,omitempty"` // Hotel Amenity Category Names joined by ampersand `&`
}

// HotelAmenityTypeAPI Hotel Amenity Type API
type HotelAmenityTypeAPI struct {
	HotelAmenityTypeCode                 *int                                      `json:"hotel_amenity_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null"`     // Hotel Amenity Type Code
	HotelAmenityTypeName                 *string                                   `json:"hotel_amenity_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Hotel Amenity Type Name
	HotelAmenityCategoryHotelAmenityType []HotelAmenityCategoryHotelAmenityTypeAPI `json:"hotel_amenity_category_hotel_amenity_type,omitempty" gorm:"-"`                                               // Hotel Amenity Category Hotel Amenity Type
	HotelAmenityTypeAsset                *HotelAmenityTypeAssetAPI                 `json:"hotel_amenity_type_asset,omitempty" gorm:"-"`                                                                // Hotel Amenity Type Asset
}
