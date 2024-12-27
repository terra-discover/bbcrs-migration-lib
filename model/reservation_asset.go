package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReservationAsset Reservation Asset
type ReservationAsset struct {
	basic.Base
	basic.DataOwner
	ReservationAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// ReservationAssetAPI Reservation Asset API
type ReservationAssetAPI struct {
	ReservationID           *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
