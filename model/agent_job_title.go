package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentJobTitle Model
type AgentJobTitle struct {
	basic.Base
	basic.DataOwner
	AgentJobTitleAPI
	JobTitle *JobTitle `json:"job_title" gorm:"foreignKey:JobTitleID;references:ID"`
	Agent    *Agent    `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentJobTitleAPI API
type AgentJobTitleAPI struct {
	AgentID    *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	JobTitleID *uuid.UUID `json:"job_title_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsDefault  *bool      `json:"is_default,omitempty"`
}
