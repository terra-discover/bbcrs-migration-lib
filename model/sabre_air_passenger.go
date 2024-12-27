package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type SabreAirPassengerTotalFareMessage struct {
	basic.Base
	basic.DataOwner
	SessionID       *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FlightID        *uuid.UUID `json:"flight_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FareMessageType *string    `json:"fare_message_type,omitempty" gorm:"type:varchar(8)"`
	FareMessageCode *string    `json:"fare_message_code,omitempty" gorm:"type:varchar(8)"`
	AirlineCode     *string    `json:"airline_code,omitempty" gorm:"type:varchar(16)"`
	Message         *string    `json:"message,omitempty" gorm:"type:text"`
}
