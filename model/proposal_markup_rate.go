package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalMarkupRate struct {
	basic.Base
	basic.DataOwner
	ProposalMarkupRateAPI
	Proposal *Proposal `json:"proposal,omitempty" gorm:"foreignKey:ProposalID;references:ID"`
}

type ProposalMarkupRateAPI struct {
	ProposalID         *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	MarkupCategoryID   *uuid.UUID `json:"markup_category_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	FeeTaxTypeID       *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	FeeTaxTypeCode     *string    `json:"fee_tax_type_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	FeeTaxTypeName     *string    `json:"fee_tax_type_name,omitempty" gorm:"type:varchar(256)" swaggertype:"string"`
	IsSummary          *bool      `json:"is_summary,omitempty"`
	IsTaxInclusive     *bool      `json:"is_tax_inclusive,omitempty"`
	Percent            *float64   `json:"percent,omitempty"`
	CurrencyID         *string    `json:"currency_id,omitempty" gorm:"type:varchar(36)"`
	Amount             *float64   `json:"amount,omitempty"`
	ChargeTypeID       *string    `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	ParentMarkupRateID *string    `json:"parent_markup_rate_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	Cascade            *bool      `json:"cascade,omitempty"`
	ApplicationOrder   *int       `json:"application_order,omitempty" gorm:"type:int"`
	TotalAmount        *float64   `json:"total_amount,omitempty"`
	Description        *string    `json:"description,omitempty" gorm:"type:text"`
}
