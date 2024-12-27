package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentCorporateGroup Agent Corporate Group
type AgentCorporateGroup struct {
	basic.Base
	basic.DataOwner
	AgentCorporateGroupAPI
	Agent *Agent `json:"agent,omitempty"`
}

// AgentCorporateGroupAPI Agent Corporate Group Api
type AgentCorporateGroupAPI struct {
	AgentID          *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CorporateGroupID *uuid.UUID `json:"corporate_group_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
