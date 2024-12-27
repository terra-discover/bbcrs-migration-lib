package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentVoucherType Model
type AgentVoucherType struct {
	basic.Base
	basic.DataOwner
	AgentVoucherTypeAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentVoucherTypeAPI API
type AgentVoucherTypeAPI struct {
	AgentID       *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	VoucherTypeID *uuid.UUID `json:"voucher_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
