package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAmenityCategoryHotelAmenityType struct
type HotelAmenityCategoryHotelAmenityType struct {
	basic.Base
	basic.DataOwner
	HotelAmenityCategoryHotelAmenityTypeAPI
	HotelAmenityCategory *HotelAmenityCategory `json:"hotel_amenity_category,omitempty"`                                                                    // Hotel Amenity Category
	HotelAmenityTypeID   *uuid.UUID            `json:"hotel_amenity_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Hotel Amenity Type ID
}

// HotelAmenityCategoryHotelAmenityTypeAPI for handle request body
type HotelAmenityCategoryHotelAmenityTypeAPI struct {
	HotelAmenityCategoryID *uuid.UUID `json:"hotel_amenity_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Hotel Amenity Category ID
}
