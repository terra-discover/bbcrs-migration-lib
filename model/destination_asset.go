package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DestinationAsset struct
type DestinationAsset struct {
	basic.Base
	basic.DataOwner
	DestinationAssetAPI
	Destination           *Destination           `json:"destination,omitempty" swaggerignore:"true"`                                                                                                                             // destination
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                                                                       // multimedia description
	DestinationID         *uuid.UUID             `json:"destination_id,omitempty" gorm:"type:varchar(36);index:destination_asset_multimedia_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // destination id
}

// DestinationAssetAPI struct writable by request body
type DestinationAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:destination_asset_multimedia_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
