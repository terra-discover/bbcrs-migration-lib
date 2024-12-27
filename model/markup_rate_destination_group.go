package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateDestinationGroup Model
type MarkupRateDestinationGroup struct {
	basic.Base
	basic.DataOwner
	MarkupRateDestinationGroupAPI
	MarkupRate       *MarkupRate       `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	DestinationGroup *DestinationGroup `json:"destination_group" gorm:"foreignKey:DestinationGroupID;references:ID"`
}

// MarkupRateDestinationGroupAPI API
type MarkupRateDestinationGroupAPI struct {
	MarkupRateID       *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded         *bool      `json:"is_included,omitempty"`
}
