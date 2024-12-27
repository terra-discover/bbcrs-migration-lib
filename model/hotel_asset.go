package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAsset Hotel Asset
type HotelAsset struct {
	basic.Base
	basic.DataOwner
	HotelAssetAPI
	HotelID               *uuid.UUID             `json:"hotel_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // Hotel ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                       // Multimedia Description
}

// HotelAssetAPI Hotel Asset API
type HotelAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // Multimedia Description Id
}
