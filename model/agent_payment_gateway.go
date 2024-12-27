package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentPaymentGateway Agent Payment Gateway
type AgentPaymentGateway struct {
	basic.Base
	basic.DataOwner
	AgentPaymentGatewayAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentPaymentGatewayAPI Agent Payment Gateway Api
type AgentPaymentGatewayAPI struct {
	AgentID          *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	PaymentGatewayID *uuid.UUID `json:"payment_gateway_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsPrimary        *bool      `json:"is_primary,omitempty" gorm:"type:bool"`
}
