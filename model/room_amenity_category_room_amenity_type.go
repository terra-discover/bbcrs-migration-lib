package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomAmenityCategoryRoomAmenityType struct
type RoomAmenityCategoryRoomAmenityType struct {
	basic.Base
	basic.DataOwner
	RoomAmenityCategoryRoomAmenityTypeAPI
	RoomAmenityTypeID   *uuid.UUID           `json:"room_amenity_type_id,omitempty" gorm:"type:varchar(36);index:idx_room_amenity_type__room_amenity_category_id,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Room Amenity Type Id
	RoomAmenityCategory *RoomAmenityCategory `json:"room_amenity_category,omitempty"`                                                                                                                                                          // Room Amenity Type Id
}

// RoomAmenityCategoryRoomAmenityTypeAPI struct
type RoomAmenityCategoryRoomAmenityTypeAPI struct {
	RoomAmenityCategoryID *uuid.UUID `json:"room_amenity_category_id,omitempty" gorm:"type:varchar(36);index:idx_room_amenity_type__room_amenity_category_id,unique,where:deleted_at is null;not null"` // Room Amenity Category Id
}
