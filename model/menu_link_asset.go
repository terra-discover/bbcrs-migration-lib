package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MenuLinkAsset struct
type MenuLinkAsset struct {
	basic.Base
	basic.DataOwner
	MenuLinkID            *uuid.UUID             `json:"menu_link_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // MenuLink ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                                 // MultimediaDescription
	MenuLinkAssetAPI
}

// MenuLinkAssetAPI struct writable by request body
type MenuLinkAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
