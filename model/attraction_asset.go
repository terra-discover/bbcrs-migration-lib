package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AttractionAsset struct
type AttractionAsset struct {
	basic.Base
	basic.DataOwner
	AttractionAssetAPI
	Attraction            *Attraction            `json:"attraction,omitempty" swaggerignore:"true"`                                                                                                                            // attraction
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                                                                     // multimedia description
	AttractionID          *uuid.UUID             `json:"attraction_id,omitempty" gorm:"type:varchar(36);index:attraction_asset_multimedia_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // attraction id
}

// AttractionAssetAPI struct writable by request body
type AttractionAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:attraction_asset_multimedia_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
