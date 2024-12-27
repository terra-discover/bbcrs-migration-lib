package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateProject struct
type CorporateProject struct {
	basic.Base
	basic.DataOwner
	CorporateProjectAPI
	Corporate *Corporate `json:"corporate,omitempty"` // Corporate
	Project   *Project   `json:"project,omitempty"`   // Project
}

// CorporateProjectAPI struct for payload
type CorporateProjectAPI struct {
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null"` // Corporate ID
	ProjectID   *uuid.UUID `json:"project_id,omitempty" gorm:"type:varchar(36);not null"`   // Project ID
}
