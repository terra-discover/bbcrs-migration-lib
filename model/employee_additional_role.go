package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EmployeeAdditionalRole EmployeeAdditionalRole
type EmployeeAdditionalRole struct {
	basic.Base
	basic.DataOwner
	EmployeeAdditionalRoleAPI
	Employee *Employee `json:"employee,omitempty"`
	JobTitle *JobTitle `json:"job_title,omitempty" gorm:"foreignKey:JobTitleID;references:ID"`
	Division *Division `json:"division,omitempty" gorm:"foreignKey:DivisionID;references:ID"`
}

// EmployeeAdditionalRoleAPI EmployeeAdditionalRole API
type EmployeeAdditionalRoleAPI struct {
	EmployeeID *uuid.UUID `json:"employee_id,omitempty" swaggertype:"string" format:"uuid" gorm:"not null" validate:"required"`
	JobTitleID *uuid.UUID `json:"job_title_id,omitempty" swaggertype:"string" format:"uuid" gorm:"not null" validate:"required"`
	DivisionID *uuid.UUID `json:"division_id,omitempty" swaggertype:"string" format:"uuid" gorm:"not null" validate:"required"`
}
