package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Proposal Payload
type ProposalPayload struct {
	basic.Base
	basic.DataOwner
	ProposalPayloadAPI
}

// ProposalPayloadAPI Proposal Payload API
type ProposalPayloadAPI struct {
	ProposalID            *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	DistributionPartnerID *uuid.UUID `json:"distribution_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	IntegrationPartnerID  *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ReferenceID           *uuid.UUID `json:"reference_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Request               *string    `json:"request,omitempty" gorm:"type:text"`
	Response              *string    `json:"response,omitempty" gorm:"type:text"`
}
