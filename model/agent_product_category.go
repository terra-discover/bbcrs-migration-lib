package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentProductCategory struct
type AgentProductCategory struct {
	basic.Base
	basic.DataOwner
	AgentProductCategoryAPI
	Agent *Agent `json:"agent,omitempty"`
	//ProductCategory *ProductCategory `json:"product_category,omitempty"`
}

// AgentProductCategoryAPI struct
type AgentProductCategoryAPI struct {
	AgentID           *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;"`            // Agent ID
	ProductCategoryID *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(36);not null;"` // Product Category ID
}
