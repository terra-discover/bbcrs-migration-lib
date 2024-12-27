package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserTypeCorporate Agent User Type Corporate
type AgentUserTypeCorporate struct {
	basic.Base
	basic.DataOwner
	AgentUserTypeCorporateAPI
	Corporate *Corporate `json:"corporate,omitempty"`
}

// AgentUserTypeCorporateAPI Agent User Type Corporate Api
type AgentUserTypeCorporateAPI struct {
	AgentUserTypeID *uuid.UUID `json:"agent_user_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	CorporateID     *uuid.UUID `json:"corporate_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsAllowed       *bool      `json:"is_allowed,omitempty"`
}
