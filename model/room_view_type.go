package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomViewType Room View Type
type RoomViewType struct {
	basic.Base
	basic.DataOwner
	RoomViewTypeAPI
	RoomViewTypeTranslation *RoomViewTypeTranslation `json:"room_view_type_translation,omitempty"`
}

// RoomViewTypeAPI Room View Type API
type RoomViewTypeAPI struct {
	RoomViewTypeCode *int    `json:"room_view_type_code,omitempty" gorm:"type:smallint;not null;index:idx_room_view_type_code_deleted_at,unique,where:deleted_at is null" example:"1"`                // Room View Type Code
	RoomViewTypeName *string `json:"room_view_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_room_view_type_name_deleted_at,unique,where:deleted_at is null" example:"Airport view"` // Room View Type Name
}

// Seed RoomViewType
func (roomViewType *RoomViewType) Seed() *RoomViewType {
	seed := RoomViewType{
		RoomViewTypeAPI: RoomViewTypeAPI{
			RoomViewTypeCode: lib.Intptr(1),
			RoomViewTypeName: lib.Strptr("Airport view"),
		},
	}
	_ = lib.Merge(seed, &roomViewType)
	return roomViewType
}
