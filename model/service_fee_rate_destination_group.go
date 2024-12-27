package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateDestinationGroup Model
type ServiceFeeRateDestinationGroup struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateDestinationGroupAPI
	ServiceFeeRate   *ServiceFeeRate   `json:"service_fee_rate" gorm:"foreignKey:ServiceFeeRateID;references:ID"`
	DestinationGroup *DestinationGroup `json:"destination_group" gorm:"foreignKey:DestinationGroupID;references:ID"`
}

// ServiceFeeRateDestinationGroupAPI API
type ServiceFeeRateDestinationGroupAPI struct {
	ServiceFeeRateID   *uuid.UUID `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_destination_group__destination_group_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_destination_group__destination_group_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	IsIncluded         *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
