package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TaxRateDestinationGroup model
type TaxRateDestinationGroup struct {
	basic.Base
	basic.DataOwner
	TaxRateDestinationGroupAPI
	DestinationGroup *DestinationGroup `json:"destination_group,omitempty"`
}

// TaxRateDestinationGroupAPI model
type TaxRateDestinationGroupAPI struct {
	TaxRateID          *uuid.UUID `json:"tax_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded         *bool      `json:"is_included,omitempty"`
}
