package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeAsset Model
type ThemeAsset struct {
	basic.Base
	basic.DataOwner
	ThemeID               *uuid.UUID             `json:"theme_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Social Network ID
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                       // multimedia description
	ThemeAssetAPI
}

// ThemeAssetAPI Model
type ThemeAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggertype:"string"` // Multimedia Description ID
}
