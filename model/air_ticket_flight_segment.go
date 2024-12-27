package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirTicketFlightSegment struct {
	basic.Base
	basic.DataOwner
	AirTicketFlightSegmentAPI
	AirTicket     *AirTicket     `json:"air_ticket,omitempty" gorm:"foreignKey:AirTicketID;references:ID"`
	FlightSegment *FlightSegment `json:"flight_segment,omitempty" gorm:"foreignKey:FlightSegmentID;references:ID"`
}

type AirTicketFlightSegmentAPI struct {
	AirTicketID     *uuid.UUID `json:"air_ticket_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`     // The reference to the air ticket.
	FlightSegmentID *uuid.UUID `json:"flight_segment_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // The reference to the flight segment.
	TicketStatus    *int       `json:"ticket_status,omitempty" gorm:"type:smallint"`                                                    // Indicates the status of the ticket, e.g. open, used, exchanged, etc.
}
