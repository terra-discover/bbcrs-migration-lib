package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserTypeSupplier Agent User Type Supplier
type AgentUserTypeSupplier struct {
	basic.Base
	basic.DataOwner
	AgentUserTypeSupplierAPI
	AgentUser    *AgentUser    `json:"agent_user,omitempty"`
	SupplierType *SupplierType `json:"supplier_type,omitempty"`
}

// AgentUserTypeSupplierAPI Agent User Type Supplier Api
type AgentUserTypeSupplierAPI struct {
	AgentUserID    *uuid.UUID `json:"agent_user_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsAllowed      *bool      `json:"is_allowed,omitempty" gorm:"type:bool"`
}
