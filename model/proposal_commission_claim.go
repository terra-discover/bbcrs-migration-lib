package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProposalCommissionClaim Proposal Commission Claim
type ProposalCommissionClaim struct {
	basic.Base
	basic.DataOwner
	ProposalCommissionClaimAPI
}

// ProposalCommissionClaimAPI Proposal Commission Claim API
type ProposalCommissionClaimAPI struct {
	ProposalID        *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // The reference to the proposal.
	CommissionClaimID *uuid.UUID `json:"commission_claim_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`  // The reference to the commission claim.
	Percent           *float64   `json:"percent,omitempty"`                                                                         // Percentage of the commission being claimed, e.g. 10 for 10% commission.
	CurrencyID        *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // The reference to the currency of the money amount.
	Amount            *float64   `json:"amount,omitempty"`                                                                          // Fixed amount of commission being claimed.
	ChargeTypeID      *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`       // The type of the amount being charged, e.g. per night.
	IsIncluded        *bool      `json:"is_included,omitempty"`                                                                     // Indicates whether this commission is included or separated from the fare.
	SharingPercent    *float64   `json:"sharing_percent,omitempty"`                                                                 // Percentage of the commission being given to another party, e.g. 3 for 3% sharing commission.
	ReferenceID       *uuid.UUID `json:"reference_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`         // A reference to an air itinerary, or air origin destination option, or flight segment, or room stay.
}
