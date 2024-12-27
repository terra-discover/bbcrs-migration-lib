package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyProgramAsset struct
type LoyaltyProgramAsset struct {
	basic.Base
	basic.DataOwner
	LoyaltyProgramAssetAPI
	LoyaltyProgram        *LoyaltyProgram        `json:"loyalty_program,omitempty" swaggerignore:"true"`
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// LoyaltyProgramAssetAPI for secure request body API
type LoyaltyProgramAssetAPI struct {
	LoyaltyProgramID        *uuid.UUID `json:"loyalty_program_id,omitempty"`
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty"`
}
