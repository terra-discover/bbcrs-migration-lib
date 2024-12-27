package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentIdentityRuleHistory Agent Identity Rule History
type AgentIdentityRuleHistory struct {
	basic.Base
	basic.DataOwner
	AgentIdentityRuleHistoryAPI
}

// AgentIdentityRuleHistoryAPI Agent Identity Rule History Api
type AgentIdentityRuleHistoryAPI struct {
	AgentIdentityRuleID *uuid.UUID `json:"agent_identity_rule_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"`
	AgentID             *uuid.UUID `json:"agent_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"`
	IdentityCode        *string    `json:"identity_code,omitempty" gorm:"type:varchar(16);not null"`
	IdentityName        *string    `json:"identity_name,omitempty" gorm:"type:varchar(256);not null"`
	Prefix              *string    `json:"prefix,omitempty" gorm:"type:varchar(16)"`
	DynamicPrefix       *string    `json:"dynamic_prefix,omitempty" gorm:"varchar(36)"`
	NextNumber          *string    `json:"next_number,omitempty" gorm:"varchar(16)"`
	IsReset             *bool      `json:"is_reset,omitempty" gorm:"type:bool"`
	ResetFrequency      *string    `json:"reset_frequency,omitempty" gorm:"type:char(1)"`
	LastNumber          *string    `json:"last_number,omitempty" gorm:"varchar(16)"`
	LastUsedCounter     *int64     `json:"last_used_counter,omitempty" gorm:"type:bigint"`
	LastYear            *int       `json:"last_year,omitempty" gorm:"type:smallint"`
	LastMonth           *int       `json:"last_month,omitempty" gorm:"type:smallint"`
	Description         *string    `json:"description,omitempty" gorm:"type:text"`
}
