package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateDestinationGroup Model
type ProcessingFeeRateDestinationGroup struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateDestinationGroupAPI
	ProcessingFeeRate *ProcessingFeeRate `json:"processing_fee_rate" gorm:"foreignKey:ProcessingFeeRateID;references:ID"`
	DestinationGroup  *DestinationGroup  `json:"destination_group" gorm:"foreignKey:DestinationGroupID;references:ID"`
}

// ProcessingFeeRateDestinationGroupAPI API
type ProcessingFeeRateDestinationGroupAPI struct {
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	DestinationGroupID  *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded          *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
