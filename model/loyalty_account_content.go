package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// LoyaltyAccountContent struct
type LoyaltyAccountContent struct {
	basic.Base
	basic.DataOwner
	LoyaltyAccountContentAPI
	LoyaltyAccount *LoyaltyAccount `json:"loyalty_account,omitempty"`
}

// LoyaltyAccountContentAPI for secure request body API
type LoyaltyAccountContentAPI struct {
	LoyaltyAccountID     *uuid.UUID `json:"loyalty_account_id,omitempty"`
	ContentDescriptionID *uuid.UUID `json:"content_description_id,omitempty"`
}
