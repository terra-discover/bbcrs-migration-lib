package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Proposal Budget
type ProposalBudget struct {
	basic.Base
	basic.DataOwner
	ProposalBudgetAPI
}

// ProposalBudgetAPI Proposal Budget API
type ProposalBudgetAPI struct {
	ProposalID   *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // The reference to the proposal.
	ProjectID    *uuid.UUID `json:"project_id,omitempty"  gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // The reference to the project.
	CostCenterID *uuid.UUID `json:"cost_center_id,omitempty"  gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`      // The reference to the cost center.
	ReferenceID  *uuid.UUID `json:"reference_id,omitempty"  gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`        // A reference to an air itinerary, or air origin destination option, or flight segment, or room stay.
}
