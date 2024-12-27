package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentTaxCategory Model
type AgentTaxCategory struct {
	basic.Base
	basic.DataOwner
	AgentTaxCategoryAPI
	Agent *Agent `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentTaxCategoryAPI API
type AgentTaxCategoryAPI struct {
	AgentID       *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	TaxCategoryID *uuid.UUID `json:"tax_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsDefault     *bool      `json:"is_default,omitempty"`
}
