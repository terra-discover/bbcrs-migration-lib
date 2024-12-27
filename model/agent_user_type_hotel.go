package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserTypeHotel Agent User Type Hotel
type AgentUserTypeHotel struct {
	basic.Base
	basic.DataOwner
	AgentUserTypeHotelAPI
	AgentUserType *AgentUserType `json:"agent_user_type,omitempty"`
	Hotel         *Hotel         `json:"hotel,omitempty"`
}

// AgentUserTypeHotelAPI Agent User Type Hotel Api
type AgentUserTypeHotelAPI struct {
	AgentUserTypeID *uuid.UUID `json:"agent_user_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	HotelID         *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IsAllowed       *bool      `json:"is_allowed,omitempty"`
}
