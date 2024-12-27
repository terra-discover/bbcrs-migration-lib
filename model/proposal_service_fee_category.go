package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalServiceFeeCategory struct {
	basic.Base
	basic.DataOwner
	ProposalServiceFeeCategoryAPI
}

type ProposalServiceFeeCategoryAPI struct {
	ProposalID             *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	ServiceFeeCategoryID   *uuid.UUID `json:"service_fee_category_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	ServiceFeeCategoryCode *string    `json:"service_fee_category_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	ServiceFeeCategoryName *string    `json:"service_fee_category_name,omitempty" gorm:"type:varchar(256)" swaggertype:"string"`
	IsIncluded             *bool      `json:"is_included,omitempty"`
	IsHidden               *bool      `json:"is_hidden,omitempty"`
	Description            *string    `json:"description,omitempty" gorm:"type:text"`
}
