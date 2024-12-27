package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserCorporate Agent User Corporate
type AgentUserCorporate struct {
	basic.Base
	basic.DataOwner
	AgentUserCorporateAPI
	AgentUser *AgentUser `json:"agent_user,omitempty"`
	Corporate *Corporate `json:"corporate,omitempty"`
}

// AgentUserCorporateAPI Agent User Corporate Api
type AgentUserCorporateAPI struct {
	AgentUserID *uuid.UUID `json:"agent_user_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsAllowed   *bool      `json:"is_allowed,omitempty"`
}
