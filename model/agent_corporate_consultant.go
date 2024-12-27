package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentCorporateConsultant Model
type AgentCorporateConsultant struct {
	basic.Base
	basic.DataOwner
	AgentCorporateConsultantAPI
	AgentCorporate *AgentCorporate `json:"agent_corporate" gorm:"foreignKey:AgentCorporateID;references:ID"`
	Employee       *Employee       `json:"employee" gorm:"foreignKey:EmployeeID;references:ID"`
}

// AgentCorporateConsultantAPI API
type AgentCorporateConsultantAPI struct {
	AgentCorporateID *uuid.UUID `json:"agent_corporate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	EmployeeID       *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
