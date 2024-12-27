package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirTicketTraveler struct {
	basic.Base
	basic.DataOwner
	AirTicketTravelerAPI
	AirTicket   *AirTicket   `json:"air_ticket,omitempty" gorm:"foreignKey:AirTicketID;references:ID"`
	AirTraveler *AirTraveler `json:"air_traveler,omitempty" gorm:"foreignKey:AirTravelerID;references:ID"`
}

type AirTicketTravelerAPI struct {
	AirTicketID   *uuid.UUID `json:"air_ticket_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AirTravelerID *uuid.UUID `json:"air_traveler_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
