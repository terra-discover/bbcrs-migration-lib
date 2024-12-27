package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SocialNetworkAsset Model
type SocialNetworkAsset struct {
	basic.Base
	basic.DataOwner
	SocialNetworkID       *uuid.UUID             `json:"social_network_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Social Network ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                // multimedia description
	SocialNetworkAssetAPI
}

// SocialNetworkAssetAPI Model
type SocialNetworkAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Multimedia Description ID
}
