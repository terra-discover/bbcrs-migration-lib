package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalTaxCategory struct {
	basic.Base
	basic.DataOwner
	ProposalTaxCategoryAPI
}

type ProposalTaxCategoryAPI struct {
	ProposalID      *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null;"`
	TaxCategoryID   *uuid.UUID `json:"tax_category_id,omitempty" gorm:"type:varchar(36);not null;"`
	TaxCategoryCode *string    `json:"tax_category_code,omitempty" gorm:"type:varchar(36)"`
	TaxCategoryName *string    `json:"tax_category_name,omitempty" gorm:"type:varchar(256)"`
	Description     *string    `json:"description,omitempty" gorm:"type:text"`
}
