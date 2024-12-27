package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReservationRoomType Reservation Room Type
type ReservationRoomType struct {
	basic.Base
	basic.DataOwner
	ReservationRoomTypeAPI
}

// ReservationRoomTypeAPI Reservation Room Type Api
type ReservationRoomTypeAPI struct {
	ReservationID *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	RoomTypeID    *uuid.UUID `json:"room_type_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null" swaggertype:"string" format:"uuid"`
	RoomTypeCode  *string    `json:"room_type_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	RoomTypeName  *string    `json:"room_type_name,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	IsAccrual     *bool      `json:"is_accrual,omitempty" gorm:"type:bool;default:false" swaggertype:"boolean"`
	IsRedeemable  *bool      `json:"is_redeemable,omitempty" gorm:"type:bool;default:false" swaggertype:"boolean"`
}
