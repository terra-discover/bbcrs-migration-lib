package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CallOutAsset struct
type CallOutAsset struct {
	basic.Base
	basic.DataOwner
	CallOutAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
	CallOutID             *uuid.UUID             `json:"call_out_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // CallOut id
}

// CallOutAssetAPI struct
type CallOutAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
