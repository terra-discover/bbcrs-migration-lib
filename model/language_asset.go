package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LanguageAsset non writable by request body
type LanguageAsset struct {
	basic.Base
	basic.DataOwner
	LanguageID            *uuid.UUID             `json:"language_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Language ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                                // multimedia description
	LanguageAssetAPI
}

// LanguageAssetAPI writable by request body
type LanguageAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // to get Multimedia Description ID, please refer this API https://lab.tog.co.id/bb/multimedia-service/-/blob/master/docs/swagger.yaml
}
