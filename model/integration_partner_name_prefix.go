package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerNamePrefix Integration Partner Name Prefix
type IntegrationPartnerNamePrefix struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerNamePrefixAPI
}

// IntegrationPartnerNamePrefixAPI Integration Partner Name Prefix API
type IntegrationPartnerNamePrefixAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`         // Integration Partner ID
	NamePrefixID         *uuid.UUID `json:"name_prefix_id,omitempty" gorm:"type:varchar(36)"`                           // Name Prefix ID
	NamePrefixCode       *string    `json:"name_prefix_code,omitempty" gorm:"type:varchar(16);not null" example:"bybn"` // Name Prefix Code
	NamePrefixName       *string    `json:"name_prefix_name,omitempty" gorm:"type:varchar(256)" example:"bayu buana"`   // Name Prefix Name
}
