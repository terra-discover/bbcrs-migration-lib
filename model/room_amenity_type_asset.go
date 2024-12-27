package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomAmenityTypeAsset struct
type RoomAmenityTypeAsset struct {
	basic.Base
	basic.DataOwner
	RoomAmenityTypeID *uuid.UUID `json:"room_amenity_type_id,omitempty" gorm:"type:varchar(36);index:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	RoomAmenityTypeAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// RoomAmenityTypeAssetAPI for handle request body
type RoomAmenityTypeAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
