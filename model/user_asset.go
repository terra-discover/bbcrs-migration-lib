package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserAsset struct
type UserAsset struct {
	basic.Base
	basic.DataOwner
	UserAssetAPI
	UserAccount           *UserAccount           `json:"user_account,omitempty" swaggerignore:"true"` // user
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`            // multimedia description
	UserAccountID         *uuid.UUID             `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"`
}

// UserAssetAPI struct writable by request body
type UserAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);index:user_asset_multimedia_unique,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
