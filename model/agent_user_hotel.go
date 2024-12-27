package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentUserHotel Agent User Hotel
type AgentUserHotel struct {
	basic.Base
	basic.DataOwner
	AgentUserHotelAPI
	AgentUser *AgentUser `json:"agent_user,omitempty"`
	Hotel     *Hotel     `json:"hotel,omitempty"`
}

// AgentUserHotelAPI Agent User Hotel Api
type AgentUserHotelAPI struct {
	AgentUserID *uuid.UUID `json:"agent_user_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	HotelID     *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IsAllowed   *bool      `json:"is_allowed,omitempty"`
}
