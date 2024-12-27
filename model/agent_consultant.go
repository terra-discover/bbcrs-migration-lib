package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentConsultant Model
type AgentConsultant struct {
	basic.Base
	basic.DataOwner
	AgentConsultantAPI
	Employee *Employee `json:"employee,omitempty" gorm:"foreignKey:EmployeeID;references:ID"`
	Agent    *Agent    `json:"agent,omitempty" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentConsultantAPI API
type AgentConsultantAPI struct {
	AgentID        *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);index:idx_agent_id__employee_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	EmployeeID     *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36);index:idx_agent_id__employee_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	CanIssueTicket *bool      `json:"can_issue_ticket,omitempty"`
}

// AgentConsultantDataGet schema
type AgentConsultantDataGet struct {
	ID             *uuid.UUID       `json:"id,omitempty"`
	CreatedAt      *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt      *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort           *int64           `json:"sort,omitempty"`                                               // sort
	JobTitleName   *string          `json:"job_title_name,omitempty"`                                     // job title name
	GivenName      *string          `json:"given_name,omitempty"`                                         // given name
	MiddleName     *string          `json:"middle_name,omitempty"`                                        // middle name
	Surname        *string          `json:"surname,omitempty"`                                            // surname
	Status         *int64           `json:"status,omitempty"`
	CanIssueTicket *bool            `json:"can_issue_ticket,omitempty"`
	AgentID        *uuid.UUID       `json:"-"`
	EmployeeID     *uuid.UUID       `json:"employee_id,omitempty"`
	TeamCount      *int             `json:"team_count,omitempty"`
	AgentConsultantDataGetAPI
}

// AgentConsultantDataGetAPI schema
type AgentConsultantDataGetAPI struct {
	Person   *Person   `json:"person,omitempty" gorm:"foreignKey:ID"`
	Office   *Office   `json:"office" gorm:"foreignKey:ID"`
	JobTitle *JobTitle `json:"job_title" gorm:"foreignKey:ID"`
}

// AgentConsultantRequestAPI API
type AgentConsultantRequestAPI struct {
	Employee *[]AgentConsultantRequest `json:"employee" gorm:"-" validate:"required"`
}

// AgentConsultantRequest schema
type AgentConsultantRequest struct {
	EmployeeID     *uuid.UUID `json:"employee_id,omitempty"  gorm:"type:varchar(36)" format:"uuid"`
	CanIssueTicket *bool      `json:"can_issue_ticket,omitempty"`
}
