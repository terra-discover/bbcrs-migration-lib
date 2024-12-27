package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionClaimDepartureDate Commission Claim Departure Date
type CommissionClaimDepartureDate struct {
	basic.Base
	basic.DataOwner
	CommissionClaimDepartureDateAPI
	CommissionClaim *CommissionClaim `json:"commission_claim,omitempty" gorm:"foreignKey:CommissionClaimID;references:ID" swaggerignore:"true"`
}

// CommissionClaimDepartureDateAPI Commission Claim Departure Date API
type CommissionClaimDepartureDateAPI struct {
	CommissionClaimID   *uuid.UUID       `json:"commision_claim_id,omitempty" gorm:"not null" swaggertype:"string" format:"uuid"`      // Commission Claim ID
	StartDate           *strfmt.DateTime `json:"start_date,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"` // Start Date
	EndDate             *strfmt.DateTime `json:"end_date,omitempty" format:"date-time" gorm:"type:timestamptz" swaggertype:"string"`   // End Date
	Monday              *bool            `json:"monday,omitempty" example:"true"`                                                      // Monday
	Tuesday             *bool            `json:"tuesday,omitempty" example:"true"`                                                     // Tuesday
	Wednesday           *bool            `json:"wednesday,omitempty" example:"true"`                                                   // Wednesday
	Thursday            *bool            `json:"thursday,omitempty" example:"true"`                                                    // Thursday
	Friday              *bool            `json:"friday,omitempty" example:"true"`                                                      // Friday
	Saturday            *bool            `json:"saturday,omitempty" example:"true"`                                                    // Saturday
	Sunday              *bool            `json:"sunday,omitempty" example:"true"`                                                      // Sunday
	DepartureDateStatus *int             `json:"departure_date_status,omitempty" example:"1"`                                          // Departure Date Status
}
