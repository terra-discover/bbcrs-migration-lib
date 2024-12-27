package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MessageAsset struct
type MessageAsset struct {
	basic.Base
	basic.DataOwner
	MessageID             *uuid.UUID             `json:"message_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"` // multimedia description
	MessageAssetAPI
}

// MessageAssetAPI struct writable by request body
type MessageAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
