package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserTypeSupplierType Agent User Type Supplier Type
type AgentUserTypeSupplierType struct {
	basic.Base
	basic.DataOwner
	AgentUserTypeSupplierTypeAPI
	AgentUserType *AgentUserType `json:"agent_user_type,omitempty"`
	SupplierType  *SupplierType  `json:"supplier_type,omitempty"`
}

// AgentUserTypeSupplierTypeAPI Agent User Type Supplier Type Api
type AgentUserTypeSupplierTypeAPI struct {
	AgentUserTypeID *uuid.UUID `json:"agent_user_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	SupplierTypeID  *uuid.UUID `json:"supplier_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsAllowed       *bool      `json:"is_allowed,omitempty"`
}
