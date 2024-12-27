package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirPassengerTotalFareMessage struct {
	basic.Base
	basic.DataOwner
	AirPassengerTotalFareMessageAPI
	AirPassengerTotalFare *AirPassengerTotalFare `json:"air_passenger_total_fare,omitempty" gorm:"foreignKey:AirPassengerTotalFareID;references:ID"`
}

type AirPassengerTotalFareMessageAPI struct {
	AirPassengerTotalFareID *uuid.UUID `json:"air_passenger_total_fare_id,omitempty" gorm:"type:varchar(36);not null"`
	FareMessageType         *string    `json:"fare_message_type,omitempty" gorm:"type:varchar(8)"`
	FareMessageCode         *string    `json:"fare_message_code,omitempty" gorm:"type:varchar(8)"`
	AirlineCode             *string    `json:"airline_code,omitempty" gorm:"type:varchar(16)"`
	Title                   *string    `json:"title,omitempty" gorm:"type:varchar(256)"`
	Message                 *string    `json:"message,omitempty" gorm:"type:text"`
	IsSummary               *bool      `json:"is_summary,omitempty"`
}
