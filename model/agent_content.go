package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentContent AgentContent
type AgentContent struct {
	basic.Base
	basic.DataOwner
	AgentContentAPI
	Agent              *Agent              `json:"agent,omitempty"`
	ContentDescription *ContentDescription `json:"content_description,omitempty"`
}

// AgentContentAPI AgentContent API
type AgentContentAPI struct {
	AgentID              *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	ContentDescriptionID *uuid.UUID `json:"content_description_id,omitempty" format:"uuid" gorm:"type:varchar(36);not null;"`
}
