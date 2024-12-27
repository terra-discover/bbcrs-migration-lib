package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionClaim Commission Claim
type CommissionClaim struct {
	basic.Base
	basic.DataOwner
	CommissionClaimAPI
	Agent                            *Agent                            `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	Currency                         *Currency                         `json:"currency,omitempty"`
	ChargeType                       *ChargeType                       `json:"charge_type,omitempty"`
	CommissionClaimOriginDestination *CommissionClaimOriginDestination `json:"commission_claim_origin_destination" gorm:"foreignKey:CommissionClaimID;references:ID" swaggerignore:"true"` // Commission Claim Original Destination
	CommissionClaimIssueDate         *CommissionClaimIssueDate         `json:"commission_claim_issue_date,omitempty" gorm:"foreignKey:CommissionClaimID;references:ID"`                    // Commission Claim Issue Date
	CommissionClaimDepartureDate     *CommissionClaimDepartureDate     `json:"commission_claim_departure_date,omitempty" gorm:"foreignKey:CommissionClaimID;references:ID"`                // Commission Claim Departure Date
	CommissionClaimAirline           *CommissionClaimAirline           `json:"commission_claim_airline,omitempty" gorm:"foreignKey:CommissionClaimID;references:ID"`
	CommissionClaimBookingClass      []CommissionClaimBookingClass     `json:"commission_claim_booking_class,omitempty"`
	CommissionClaimCabinType         []CommissionClaimCabinType        `json:"commission_claim_cabin_type,omitempty"`
}

// CommissionClaimAPI Commission Claim API
type CommissionClaimAPI struct {
	AgentID             *uuid.UUID `json:"agent_id,omitempty" gorm:"not null" swaggertype:"string" format:"uuid"`      // Agent ID
	CommissionClaimCode *string    `json:"commission_claim_code,omitempty" example:"code_123" gorm:"type:varchar(36)"` // Commission Claim Code
	CommissionClaimName *string    `json:"commission_claim_name,omitempty" example:"code 123" gorm:"type:varchar(64)"` // Commission Claim Name
	Percent             *float64   `json:"percent,omitempty" gorm:"type:decimal(19,4)" example:"10.0"`                 // Percent
	CurrencyID          *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" format:"uuid"`                   // Currency ID
	Amount              *float64   `json:"amount,omitempty" gorm:"type:decimal(19,4)" example:"1000.0"`                // Amount
	ChargeTypeID        *uuid.UUID `json:"charge_type_id,omitempty" swaggertype:"string" format:"uuid"`                // Charge Type ID
	IsIncluded          *bool      `json:"is_included,omitempty" example:"true"`                                       // Is Included
	Comment             *string    `json:"comment,omitempty" example:"comment" gorm:"type:text"`                       // Comment
	Description         *string    `json:"description,omitempty" example:"desc" gorm:"type:text"`                      // Description
}
