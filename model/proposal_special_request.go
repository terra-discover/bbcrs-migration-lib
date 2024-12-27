package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalSpecialRequest struct {
	basic.Base
	basic.DataOwner
	ProposalSpecialRequestAPI
}

type ProposalSpecialRequestAPI struct {
	ProposalID         *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	SpecialRequestID   *uuid.UUID `json:"special_request_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	SpecialRequestCode *string    `json:"special_request_code,omitempty" gorm:"type:varchar(36)"`
	SpecialRequestName *string    `json:"special_request_name,omitempty" gorm:"type:varchar(256)"`
}
