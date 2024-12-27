package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Flight Violation Reason
type FlightViolationReason struct {
	basic.Base
	basic.DataOwner
	FlightViolationReasonAPI
}

// FlightViolationReasonAPI Flight Violation Reason API
type FlightViolationReasonAPI struct {
	AirOriginDestinationOptionID *uuid.UUID `json:"air_origin_destination_option_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"`
	ViolationReasonID            *uuid.UUID `json:"violation_reason_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"`
}
