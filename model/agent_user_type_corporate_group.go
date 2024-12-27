package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserTypeCorporateGroup Agent User Type Corporate Group
type AgentUserTypeCorporateGroup struct {
	basic.Base
	basic.DataOwner
	AgentUserTypeCorporateGroupAPI
}

// AgentUserTypeCorporateGroupAPI Agent User Type Corporate Group Api
type AgentUserTypeCorporateGroupAPI struct {
	AgentUserTypeID  *uuid.UUID `json:"agent_user_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	CorporateGroupID *uuid.UUID `json:"corporate_group_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsAllowed        *bool      `json:"is_allowed,omitempty"`
}
