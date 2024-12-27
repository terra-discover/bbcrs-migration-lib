package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type FlightCachingCancellationOffer struct {
	basic.Base
	basic.DataOwner
	SessionID    *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);not null;index:idx_flight_caching_cancellation_offer_session_id" swaggertype:"string"`
	ActionOffer  *string    `json:"action_offer,omitempty" gorm:"type:varchar(100)"` // In SQ partner: void_booking, void_ticket, refund_ticket
	PartnerPrice *float64   `json:"partner_price,omitempty"`                         // ex: refund price from partner
	TotalPrice   *float64   `json:"total_price,omitempty"`                           // ex: refund price from agent (partner price - refund fee)
	PenaltyFee   *float64   `json:"penalty_fee,omitempty"`                           // ex: refund fee from agent
}
