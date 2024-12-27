package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Proposal Product Purchase
type ProposalProductPurchase struct {
	basic.Base
	basic.DataOwner
	ProposalProductPurchaseAPI
	ProductPurchase *ProductPurchase `json:"product_purchase,omitempty" gorm:"foreignKey:ProductPurchaseID;references:ID"` // ProductPurchase
	Proposal        *Proposal        `json:"proposal,omitempty" gorm:"foreignKey:ProposalID;references:ID"`                // Proposal
}

// ProposalProductPurchaseAPI Proposal Product Purchase API
type ProposalProductPurchaseAPI struct {
	ProposalID        *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProductPurchaseID *uuid.UUID `json:"product_purchase_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
