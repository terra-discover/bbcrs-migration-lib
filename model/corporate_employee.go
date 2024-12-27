package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateEmployee Model
type CorporateEmployee struct {
	basic.Base
	basic.DataOwner
	CorporateEmployeeAPI
	Corporate *Corporate `json:"corporate,omitempty"`
	Employee  *Employee  `json:"employee,omitempty"`
}

// CorporateEmployeeAPI Model
type CorporateEmployeeAPI struct {
	CorporateID    *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	EmployeeID     *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`  // Employee ID
	IsTravelBooker *bool      `json:"is_travel_booker,omitempty"`
}
