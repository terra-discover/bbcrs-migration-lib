package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomLocationType Room Location Type
type RoomLocationType struct {
	basic.Base
	basic.DataOwner
	RoomLocationTypeAPI
	RoomLocationTypeTranslation *RoomLocationTypeTranslation `json:"room_location_type_translation,omitempty"`
}

// RoomLocationTypeAPI Room Location Type API
type RoomLocationTypeAPI struct {
	RoomLocationTypeCode *int    `json:"room_location_type_code,omitempty" gorm:"type:smallint;not null;index:idx_room_location_type_code_deleted_at,unique,where:deleted_at is null" example:"1"`                          // Room Location Type Code
	RoomLocationTypeName *string `json:"room_location_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_room_location_type_code_deleted_at,unique,where:deleted_at is null" example:"Away from the elevator"` // Room Location Type Name
}

// Seed RoomLocationType
func (roomLocationType *RoomLocationType) Seed() *RoomLocationType {
	seed := RoomLocationType{
		RoomLocationTypeAPI: RoomLocationTypeAPI{
			RoomLocationTypeCode: lib.Intptr(1),
			RoomLocationTypeName: lib.Strptr("Away from the elevator"),
		},
	}
	_ = lib.Merge(seed, &roomLocationType)
	return roomLocationType
}
