package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// AgentEmailTemplate Model
type AgentEmailTemplate struct {
	basic.Base
	basic.DataOwner
	AgentEmailTemplateAPI
	AgentID *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	Agent   *Agent     `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentEmailTemplateAPI API
type AgentEmailTemplateAPI struct {
	EmailTemplateID *uuid.UUID `json:"email_template_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}

// Seed AgentEmailTemplate Data
func (agentEmailTemplate *AgentEmailTemplate) Seed() *AgentEmailTemplate {
	agentID, _ := uuid.Parse("8d5fb953-73af-4a07-b450-f8641b40e383")
	emailTemplateID, _ := uuid.Parse("4fa75f64-5777-4562-b3fc-2c963f66ab00")
	agentEmailTemplateID, _ := uuid.Parse("4ba85f64-5717-4562-b3fc-2c963f66ab00")
	seed := AgentEmailTemplate{
		AgentID: &agentID,
		Base: basic.Base{
			ID: &agentEmailTemplateID,
		},
		AgentEmailTemplateAPI: AgentEmailTemplateAPI{
			EmailTemplateID: &emailTemplateID,
		},
	}
	_ = lib.Merge(seed, &agentEmailTemplate)
	return agentEmailTemplate
}
