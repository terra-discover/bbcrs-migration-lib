package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirTravelerCriteria struct {
	basic.Base
	ProposalID       *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null"`
	SeatsRequested   *int       `json:"seats_requested,omitempty" gorm:"type:smallint"`
	NumberOfAdults   *int       `json:"number_of_adults,omitempty" gorm:"type:smallint"`
	NumberOfChildren *int       `json:"number_of_children,omitempty" gorm:"type:smallint"`
	NumberOfInfants  *int       `json:"number_of_infants,omitempty" gorm:"type:smallint"`
	Proposal         *Proposal  `json:"proposal,omitempty"`
}
