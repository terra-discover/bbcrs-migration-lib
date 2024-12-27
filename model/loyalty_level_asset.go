package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyLevelAsset Loyalty Level Asset
type LoyaltyLevelAsset struct {
	basic.Base
	basic.DataOwner
	LoyaltyLevelAssetAPI
	LoyaltyLevelID        *uuid.UUID             `json:"loyalty_level_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Level Id
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                               // Multimedia Description
}

// LoyaltyLevelAssetAPI Loyalty Level Asset Api
type LoyaltyLevelAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Multimedia Description Id
}
