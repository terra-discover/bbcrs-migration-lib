package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyAccountAsset struct
type LoyaltyAccountAsset struct {
	basic.Base
	basic.DataOwner
	LoyaltyAccountAssetAPI
	LoyaltyAccount        *LoyaltyAccount        `json:"loyalty_account,omitempty"`
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`
}

// LoyaltyAccountAssetAPI for secure request body API
type LoyaltyAccountAssetAPI struct {
	LoyaltyAccountID        *uuid.UUID `json:"loyalty_account_id,omitempty"`
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty"`
}
