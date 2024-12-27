package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProductAsset struct
type ProductAsset struct {
	basic.Base
	basic.DataOwner
	ProductID             *uuid.UUID             `json:"product_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"` // multimedia description
	ProductAssetAPI
}

// ProductAssetAPI struct writable by request body
type ProductAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
