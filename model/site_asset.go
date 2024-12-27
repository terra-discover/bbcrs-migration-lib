package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteAsset struct
type SiteAsset struct {
	basic.Base
	basic.DataOwner
	SiteAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
	SiteID                *uuid.UUID             `json:"site_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Site id
}

// SiteAssetAPI struct
type SiteAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
