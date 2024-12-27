package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReservationRoomStay model
type ReservationRoomStay struct {
	basic.Base
	basic.DataOwner
	ReservationRoomStayAPI
	Reservation *Reservation `json:"reservation"`
	RoomStay    *RoomStay    `json:"room_stay"`
}

// ReservationRoomStayAPI struct
type ReservationRoomStayAPI struct {
	ReservationID *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertyp:"string" format:"uuid"` // Reservation ID
	RoomStayID    *uuid.UUID `json:"room_stay_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`  // Room Stay ID
}
