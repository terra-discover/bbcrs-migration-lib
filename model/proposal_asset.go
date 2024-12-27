package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProposalAsset Proposal Asset
type ProposalAsset struct {
	basic.Base
	basic.DataOwner
	ProposalAssetAPI
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// ProposalAssetAPI Proposal Asset API
type ProposalAssetAPI struct {
	ProposalID              *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
