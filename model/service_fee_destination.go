package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeDestination model
type ServiceFeeDestination struct {
	basic.Base
	basic.DataOwner
	ServiceFeeDestinationAPI
}

// ServiceFeeDestinationAPI model
type ServiceFeeDestinationAPI struct {
	ServiceFeeRateID   *uuid.UUID `json:"service_fee_rate_id,omitempty"`
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty"`
}
