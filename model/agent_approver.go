package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentApprover Model
type AgentApprover struct {
	basic.Base
	basic.DataOwner
	AgentApproverAPI
	Employee *Employee `json:"employee" gorm:"foreignKey:EmployeeID;references:ID"`
	Agent    *Agent    `json:"agent,omitempty" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentApproverAPI API
type AgentApproverAPI struct {
	AgentID    *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;index:idx_agent_approver__employee_id,unique,where:deleted_at is null" format:"uuid"`
	EmployeeID *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36);not null;index:idx_agent_approver__employee_id,unique,where:deleted_at is null" format:"uuid"`
}

// AgentApproverDataGet schema
type AgentApproverDataGet struct {
	ID         *uuid.UUID       `json:"id,omitempty"`
	CreatedAt  *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt  *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort       *int64           `json:"sort,omitempty"`                                               // sort
	Status     *int64           `json:"status,omitempty"`
	AgentID    *uuid.UUID       `json:"-"`
	EmployeeID *uuid.UUID       `json:"employee_id"`
	AgentApproverDataGetAPI
}

// AgentApproverDataGetAPI schema
type AgentApproverDataGetAPI struct {
	Employee *Employee `json:"employee,omitempty"`
	Person   *Person   `json:"person,omitempty" gorm:"foreignKey:ID"`
	Office   *Office   `json:"office,omitempty" gorm:"foreignKey:ID"`
}

// AgentApproverRequestAPI API
type AgentApproverRequestAPI struct {
	Employee *[]AgentApproverRequest `json:"employee" gorm:"-"`
}

// AgentApproverRequest schema
type AgentApproverRequest struct {
	EmployeeID *uuid.UUID `json:"employee_id,omitempty"  gorm:"type:varchar(36)" format:"uuid"`
}
