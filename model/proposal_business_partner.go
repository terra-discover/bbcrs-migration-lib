package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProposalBusinessPartner struct {
	basic.Base
	basic.DataOwner
	ProposalBusinessPartnerAPI
}

type ProposalBusinessPartnerAPI struct {
	ProposalID        *uuid.UUID `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	BusinessPartnerID *uuid.UUID `json:"business_partner_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
