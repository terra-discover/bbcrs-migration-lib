package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Reservation Payload
type ReservationPayload struct {
	basic.Base
	basic.DataOwner
	ProposalPayloadAPI
}

// ReservationPayloadAPI Reservation Payload API
type ReservationPayloadAPI struct {
	ReservationID         *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	DistributionPartnerID *uuid.UUID `json:"distribution_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	IntegrationPartnerID  *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ReferenceID           *uuid.UUID `json:"reference_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Request               *string    `json:"request,omitempty" gorm:"type:text"`
	Response              *string    `json:"response,omitempty" gorm:"type:text"`
}
