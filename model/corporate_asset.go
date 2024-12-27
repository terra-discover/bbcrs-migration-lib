package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateAsset Model
type CorporateAsset struct {
	basic.Base
	basic.DataOwner
	CorporateID           *uuid.UUID             `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                           // multimedia description
	CorporateAssetAPI
}

// CorporateAssetAPI Model
type CorporateAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
