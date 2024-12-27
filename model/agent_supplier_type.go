package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentSupplierType Agent Supplier Type
type AgentSupplierType struct {
	basic.Base
	basic.DataOwner
	AgentSupplierTypeAPI
	Agent        *Agent        `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
	SupplierType *SupplierType `json:"supplier_type" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentSupplierTypeAPI Agent Supplier Type Api
type AgentSupplierTypeAPI struct {
	AgentID        *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
