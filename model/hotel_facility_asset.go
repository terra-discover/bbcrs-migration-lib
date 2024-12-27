package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelFacilityAsset Hotel Facility Asset
type HotelFacilityAsset struct {
	basic.Base
	basic.DataOwner
	HotelFacilityAssetAPI
	HotelFacilityID       *uuid.UUID             `json:"hotel_facility_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Hotel Facility Id
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                // Multimedia Description
}

// HotelFacilityAssetAPI Hotel Facility Asset API
type HotelFacilityAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Multimedia Description Id
}
