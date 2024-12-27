package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentProduct struct
type AgentProduct struct {
	basic.Base
	basic.DataOwner
	AgentProductAPI
	Agent   *Agent   `json:"agent,omitempty"`
	Product *Product `json:"product,omitempty"`
}

// AgentProductAPI struct
type AgentProductAPI struct {
	AgentID   *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;"`   // Agent ID
	ProductID *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);not null;"` // Product ID
}
