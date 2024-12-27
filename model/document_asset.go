package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DocumentAsset struct
type DocumentAsset struct {
	basic.Base
	basic.DataOwner
	DocumentAssetAPI
	Document              *Document              `json:"document,omitempty"`
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// DocumentAssetAPI for secure request body API
type DocumentAssetAPI struct {
	DocumentID              *uuid.UUID `json:"document_id,omitempty" gorm:"type:varchar(36);not null"`
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null"`
}
