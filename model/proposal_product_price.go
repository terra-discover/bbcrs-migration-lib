package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Proposal Product Price
type ProposalProductPrice struct {
	basic.Base
	basic.DataOwner
	ProposalProductPriceAPI
}

// ProposalProductPriceAPI Proposal Product Price API
type ProposalProductPriceAPI struct {
	ProposalProductID          *uuid.UUID       `json:"proposal_product_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProposalProductPriceTypeID *uuid.UUID       `json:"proposal_product_price_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	StartDate                  *strfmt.DateTime `json:"start_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	EndDate                    *strfmt.DateTime `json:"end_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	Amount                     *float64         `json:"amount,omitempty"`
	AmountBeforeTax            *float64         `json:"amount_before_tax,omitempty"`
	AmountAfterTax             *float64         `json:"amount_after_tax,omitempty"`
	TotalAmountBeforeTax       *float64         `json:"total_amount_before_tax,omitempty"`
	TotalAmountAfterTax        *float64         `json:"total_amount_after_tax,omitempty"`
}
