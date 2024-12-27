package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReservationBudget Reservation Budget
type ReservationBudget struct {
	basic.Base
	basic.DataOwner
	ReservationBudgetAPI
	Project    *Project    `json:"project,omitempty"`
	CostCenter *CostCenter `json:"cost_center,omitempty"`
}

// ReservationBudgetAPI Reservation Budget API
type ReservationBudgetAPI struct {
	ReservationID *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // The reference to the reservation.
	ProjectID     *uuid.UUID `json:"project_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // The reference to the project.
	CostCenterID  *uuid.UUID `json:"cost_center_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // The reference to the cost center.
	ReferenceID   *uuid.UUID `json:"reference_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // A reference to an air itinerary, or air origin destination option, or flight segment, or room stay.
}
