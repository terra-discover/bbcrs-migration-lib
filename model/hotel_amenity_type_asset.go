package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAmenityTypeAsset struct
type HotelAmenityTypeAsset struct {
	basic.Base
	basic.DataOwner
	HotelAmenityTypeAssetAPI
	HotelAmenityTypeID    *uuid.UUID             `json:"hotel_amenity_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // Hotel Amenity Type Id
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                    // Multimedia Description
}

// HotelAmenityTypeAssetAPI struct
type HotelAmenityTypeAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // Multimedia Description Id
}
