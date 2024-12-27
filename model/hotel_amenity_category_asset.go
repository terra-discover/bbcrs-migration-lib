package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAmenityCategoryAsset struct
type HotelAmenityCategoryAsset struct {
	basic.Base
	basic.DataOwner
	HotelAmenityCategoryID *uuid.UUID             `json:"hotel_amenity_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:hotel_amenity_category_id,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	MultimediaDescription  *MultimediaDescription `json:"multimedia_description,omitempty"` // multimedia description
	HotelAmenityCategoryAssetAPI
}

// HotelAmenityCategoryAssetAPI for handle request body
type HotelAmenityCategoryAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
