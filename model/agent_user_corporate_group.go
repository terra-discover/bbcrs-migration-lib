package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserCorporateGroup Agent User Corporate Group
type AgentUserCorporateGroup struct {
	basic.Base
	basic.DataOwner
	AgentUserCorporateAPI
	AgentUser *AgentUser `json:"agent_user,omitempty"`
}

// AgentUserCorporateGroupAPI Agent User Corporate Group Api
type AgentUserCorporateGroupAPI struct {
	AgentUserID      *uuid.UUID `json:"agent_user_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	CorporateGroupID *uuid.UUID `json:"corporate_group_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsAllowed        *bool      `json:"is_allowed,omitempty"`
}
