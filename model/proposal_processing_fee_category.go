package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalProcessingFeeCategory struct {
	basic.Base
	basic.DataOwner
	ProposalProcessingFeeCategoryAPI
	Proposal *Proposal `json:"proposal,omitempty" gorm:"foreignKey:ProposalID;references:ID"`
}

type ProposalProcessingFeeCategoryAPI struct {
	ProposalID                *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	ProcessingFeeCategoryID   *uuid.UUID `json:"processing_fee_category_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	ProcessingFeeCategoryCode *string    `json:"processing_fee_category_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	ProcessingFeeCategoryName *string    `json:"processing_fee_category_name,omitempty" gorm:"type:varchar(256)" swaggertype:"string"`
	Description               *string    `json:"description,omitempty" gorm:"type:text"`
}
