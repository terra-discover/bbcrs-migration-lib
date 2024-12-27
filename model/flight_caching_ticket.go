package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingTicket
type FlightCachingTicket struct {
	basic.Base
	basic.DataOwner
	SessionID   *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_ticket_session_id;not null"`
	AirTicketID *uuid.UUID `json:"air_ticket_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	AirTicket   *AirTicket `json:"air_ticket,omitempty" gorm:"foreignKey:AirTicketID;references:ID"`
}
