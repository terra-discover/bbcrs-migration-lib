package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingCommissionClaim struct {
	basic.Base
	basic.DataOwner
	SessionID         *uuid.UUID `json:"session_id" gorm:"type:varchar(36);index:idx_flight_caching_commision_claim_session_id"`
	SharingPercent    *float64   `json:"sharing_percent,omitempty"`
	CommissionClaimID *uuid.UUID `json:"commission_claim_id" gorm:"type:varchar(36)"`
	CommissionClaimAPI
}
