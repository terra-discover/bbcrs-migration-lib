package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProposalAdditionalInfo  Proposal Additional Info
type ProposalAdditionalInfo struct {
	basic.Base
	basic.DataOwner
	ProposalAdditionalInfoAPI
	ProposalAdditionalInfoSelection []ProposalAdditionalInfoSelection `json:"proposal_additional_info_selection,omitempty"`
	AdditionalInfo                  *AdditionalInfo                   `json:"additional_info,omitempty"`
}

// ProposalAdditionalInfoAPI Proposal Additional Info API
type ProposalAdditionalInfoAPI struct {
	ProposalID          *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AdditionalInfoID    *uuid.UUID `json:"additional_info_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AdditionalInfoName  *string    `json:"additional_info_name,omitempty" gorm:"type:varchar(256)"`
	AdditionalInfoValue *string    `json:"additional_info_value,omitempty" gorm:"type:text"`
}
