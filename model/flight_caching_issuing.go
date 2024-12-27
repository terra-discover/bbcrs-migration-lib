package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingIssuing
type FlightCachingIssuing struct {
	basic.Base
	basic.DataOwner
	SessionID            *uuid.UUID       `json:"session_id,omitempty" gorm:"type:varchar(36);not null"`
	IssuingAgentID       *uuid.UUID       `json:"issuing_agent_id,omitempty"`
	IssuingAirlineCode   *string          `json:"issuing_airline_code,omitempty" example:"SQ"`
	IssuingLocation      *string          `json:"issuing_location,omitempty" example:"SQ"`
	IssuingDate          *strfmt.DateTime `json:"issuing_date,omitempty"`
	OfferExpirationDate  *strfmt.DateTime `json:"offer_expiration_date,omitempty"`
	PaymentTimeLimitDate *strfmt.DateTime `json:"payment_time_limit_date,omitempty"`
}
