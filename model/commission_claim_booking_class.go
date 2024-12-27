package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionClaim Commission Claim
type CommissionClaimBookingClass struct {
	basic.Base
	basic.DataOwner
	CommissionClaimBookingClassAPI
}

// CommissionClaimAPI Commission Claim API
type CommissionClaimBookingClassAPI struct {
	CommissionClaimID *uuid.UUID `json:"commission_claim_id,omitempty" format:"uuid" gorm:"type:varchar(36)"` // Commission Claim ID
	BookingClass      *string    `json:"booking_class,omitempty" example:"Y" gorm:"type:varchar(2)"`          // booking class
}
