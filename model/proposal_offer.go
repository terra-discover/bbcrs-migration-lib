package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProposalOffer  Proposal Offer
type ProposalOffer struct {
	basic.Base
	basic.DataOwner
	ProposalOfferAPI
}

// ProposalOfferAPI Proposal Offer API
type ProposalOfferAPI struct {
	ProposalID         *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	OfferID            *uuid.UUID `json:"offer_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	OfferCode          *string    `json:"offer_code,omitempty" gorm:"type:varchar(36)"`
	OfferName          *string    `json:"offer_name,omitempty" gorm:"type:varchar(64)"`
	OfferTypeID        *uuid.UUID `json:"offer_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ApplicationOrder   *int       `json:"application_order,omitempty" gorm:"type:int"`
	AssignAllRoomTypes *bool      `json:"assign_all_room_types,omitempty"`
	AssignAllRatePlans *bool      `json:"assign_all_rate_plans,omitempty"`
	AssignAllProducts  *bool      `json:"assign_all_products,omitempty"`
}
