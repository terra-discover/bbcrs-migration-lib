package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DeviceAsset struct
type DeviceAsset struct {
	basic.Base
	basic.DataOwner
	DeviceAssetAPI
	DeviceID              *uuid.UUID             `json:"device_id,omitempty" gorm:"type:varchar(36);index:idx_device_asset__device_id,unique,where:deleted_at is null;not null"` // Device Id
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                       // Multimedia Description
}

// DeviceAssetAPI struct
type DeviceAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Multimedia Description Id
}
