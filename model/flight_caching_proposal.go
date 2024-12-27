package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

type FlightCachingProposal struct {
	basic.Base
	basic.DataOwner
	SessionID    *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_proposal_session_id"`
	ProposalID   *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_proposal_proposal_id"`
	ProposalCode *string    `json:"proposal_code,omitempty" gorm:"type:varchar(100)"`
}

// AfterCreate Data
func (f *FlightCachingProposal) AfterCreate(tx *gorm.DB) error {
	return afterCreateFlightCachingProposal(tx, *f)
}
