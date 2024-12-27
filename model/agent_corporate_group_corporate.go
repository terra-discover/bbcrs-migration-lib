package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentCorporateGroupCorporate Agent Corporate Group Corporate
type AgentCorporateGroupCorporate struct {
	basic.Base
	basic.DataOwner
	AgentCorporateGroupCorporateAPI
	AgentCorporateGroup *AgentCorporateGroup `json:"agent_corporate_group,omitempty"`
	Corporate           *Corporate           `json:"corporate,omitempty"`
}

// AgentCorporateGroupCorporateAPI Agent Corporate Group Corporate Api
type AgentCorporateGroupCorporateAPI struct {
	AgentCorporateGroupID *uuid.UUID `json:"agent_corporate_group_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CorporateID           *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IsPrimary             *bool      `json:"is_primary,omitempty" gorm:"type:bool"`
}
