package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalAdditionalInfoSelection struct {
	basic.Base
	basic.DataOwner
	ProposalAdditionalInfoSelectionAPI
	ProposalAdditionalInfo *ProposalAdditionalInfo `json:"proposal_additional_info,omitempty" gorm:"foreignKey:ProposalAdditionalInfoID;references:ID"`
}

type ProposalAdditionalInfoSelectionAPI struct {
	ProposalAdditionalInfoID     *uuid.UUID `json:"proposal_additional_info_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AdditionalInfoOptionID       *uuid.UUID `json:"additional_info_option_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string"`
	AdditionalInfoOptionName     *string    `json:"additional_info_option_name,omitempty" gorm:"type:varchar(256)"`
	AdditionalInfoSelectionValue *string    `json:"additional_info_selection_value,omitempty" gorm:"type:text"`
}
