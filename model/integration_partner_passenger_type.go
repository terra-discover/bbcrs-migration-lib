package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerPassengerType Integration Partner Passenger Type
type IntegrationPartnerPassengerType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerPassengerTypeAPI
}

// IntegrationPartnerPassengerTypeAPI Integration Partner Passenger Type API
type IntegrationPartnerPassengerTypeAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`         // Integration Partner ID
	PassengerTypeID      *uuid.UUID `json:"passenger_type_id,omitempty" gorm:"type:varchar(36);not null;"`              // Passenger Type ID
	PassengerTypeCode    *string    `json:"passenger_type_code,omitempty" gorm:"type:varchar(36);not null" example:"1"` // Passenger Type Code
	PassengerTypeName    *string    `json:"passenger_type_name,omitempty" gorm:"type:varchar(256)" example:"Adult"`     // Passenger Type Name
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                         // Comment
}
