package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AirOriginDestinationCriteria Air Origin Destination Criteria
type AirOriginDestinationCriteria struct {
	basic.Base
	basic.DataOwner
	AirOriginDestinationCriteriaAPI
}

// AirOriginDestinationCriteriaAPI Air Origin Destination Criteria API
type AirOriginDestinationCriteriaAPI struct {
	ProposalID          *uuid.UUID       `json:"proposal_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IndexNumber         *int             `json:"index,omitempty"`
	DepartureDatetime   *strfmt.DateTime `json:"departure_datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	ArrivalDatetime     *strfmt.DateTime `json:"arrival_datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`
	OriginLocation      *string          `json:"origin,omitempty" gorm:"type:varchar(36)"`
	DestinationLocation *string          `json:"destination,omitempty" gorm:"type:varchar(36)"`
}
