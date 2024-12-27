package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateCostCenter struct
type CorporateCostCenter struct {
	basic.Base
	basic.DataOwner
	CorporateCostCenterAPI
	Corporate  *Corporate  `json:"corporate,omitempty"`
	CostCenter *CostCenter `json:"cost_center,omitempty"`
}

// CorporateCostCenterAPI for payload
type CorporateCostCenterAPI struct {
	CorporateID  *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null"`   // Corporate ID
	CostCenterID *uuid.UUID `json:"cost_center_id,omitempty" gorm:"type:varchar(36);not null"` // Cost Center ID
}
