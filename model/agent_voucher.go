package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentVoucher Model
type AgentVoucher struct {
	basic.Base
	basic.DataOwner
	AgentVoucherAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentVoucherAPI API
type AgentVoucherAPI struct {
	AgentID   *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	VoucherID *uuid.UUID `json:"voucher_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
