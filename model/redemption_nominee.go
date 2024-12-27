package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RedemptionNominee Redemption Nominee
type RedemptionNominee struct {
	basic.Base
	basic.DataOwner
	RedemptionNomineeAPI
}

// RedemptionNomineeAPI Redemption Nominee API
type RedemptionNomineeAPI struct {
	LoyaltyAccountID *uuid.UUID `json:"loyalty_account_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Loyalty Account Id
	NomineeID        *uuid.UUID `json:"nominee_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                  // Nominee Id
}
