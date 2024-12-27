package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserSupplierType Agent User Supplier Type
type AgentUserSupplierType struct {
	basic.Base
	basic.DataOwner
	AgentUserSupplierTypeAPI
	AgentUser    *AgentUser    `json:"agent_user,omitempty"`
	SupplierType *SupplierType `json:"supplier_type,omitempty"`
}

// AgentUserSupplierTypeAPI Agent User Supplier Type Api
type AgentUserSupplierTypeAPI struct {
	AgentUserID    *uuid.UUID `json:"agent_user_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IsAllowed      *bool      `json:"is_allowed,omitempty" gorm:"type:bool"`
}
