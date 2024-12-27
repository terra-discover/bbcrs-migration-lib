package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionClaim Commission Claim
type CommissionClaimCabinType struct {
	basic.Base
	basic.DataOwner
	CommissionClaimCabinTypeAPI
}

// CommissionClaimAPI Commission Claim API
type CommissionClaimCabinTypeAPI struct {
	CommissionClaimID *uuid.UUID `json:"commission_claim_id,omitempty" format:"uuid" gorm:"type:varchar(36)"` // Commission Claim ID
	CabinTypeID       *uuid.UUID `json:"cabin_type_id,omitempty" format:"uuid" gorm:"type:varchar(36)"`       // booking class
}
