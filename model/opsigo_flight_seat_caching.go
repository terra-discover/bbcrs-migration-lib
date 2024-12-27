package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type OpsigoFlightSeatCaching struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);not null" validate:"required" format:"uuid"`
	PersonID  *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36);not null" validate:"required" format:"uuid"`
	FareID    *uuid.UUID `json:"fare_id,omitempty" gorm:"type:varchar(36);not null" validate:"required" format:"uuid"`
	RouteID   *uuid.UUID `json:"route_id,omitempty" gorm:"type:varchar(36);not null" validate:"required" format:"uuid"`
	SeatID    *string    `json:"seat_id,omitempty" gorm:"type:varchar(100)" example:"PRICE3-SEG6" validate:"required"`
	Column    *string    `json:"column,omitempty" gorm:"type:varchar(100)" validate:"required" example:"A"`
	Row       *string    `json:"row,omitempty" gorm:"type:varchar(100)" validate:"required" example:"50"`
	Price     *float64   `json:"price,omitempty"`
	SeatData  *string    `json:"seat_data,omitempty"`
}
