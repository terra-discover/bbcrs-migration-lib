package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type Inquiry struct {
	basic.Base
	basic.DataOwner
	InquiryAPI
}

type InquiryAPI struct {
	AgentID          *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	CorporateID      *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	InquiryStatus    *int       `json:"inquiry_status,omitempty" gorm:"type:smallint"`
	Inquiry          *string    `json:"inquiry,omitempty" gorm:"type:text"`
	InquiryFileName  *string    `json:"inquiry_file_name,omitempty" gorm:"type:varchar(256)"`
	ReferralSourceID *string    `json:"referral_source_id,omitempty" gorm:"type:varchar(36)"`
}
