package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomAmenityCategoryAsset struct
type RoomAmenityCategoryAsset struct {
	basic.Base
	basic.DataOwner
	RoomAmenityCategoryID *uuid.UUID `json:"room_amenity_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	RoomAmenityCategoryAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// RoomAmenityCategoryAssetAPI for handle request body
type RoomAmenityCategoryAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
