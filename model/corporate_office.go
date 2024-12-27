package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateOffice Model
type CorporateOffice struct {
	basic.Base
	basic.DataOwner
	CorporateOfficeAPI
	Corporate *Corporate `json:"corporate,omitempty"`
	Office    *Office    `json:"office,omitempty"`
}

// CorporateOfficeAPI Model
type CorporateOfficeAPI struct {
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	OfficeID    *uuid.UUID `json:"office_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`    // Office ID
	IsPrimary   *bool      `json:"is_primary,omitempty"`
}
