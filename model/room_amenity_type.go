package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// RoomAmenityType Room Amenity Type
type RoomAmenityType struct {
	basic.Base
	basic.DataOwner
	RoomAmenityTypeAPI
	RoomAmenityTypeTranslation         *RoomAmenityTypeTranslation          `json:"room_amenity_type_translation,omitempty"`           // Room Amenity Type Translation
	RoomAmenityTypeAsset               *RoomAmenityTypeAsset                `json:"room_amenity_type_asset,omitempty"`                 // Room Amenity Asset
	RoomAmenityCategoryRoomAmenityType []RoomAmenityCategoryRoomAmenityType `json:"room_amenity_category_room_amenity_type,omitempty"` // Room Amenity Category Room Amenity Type
}

// RoomAmenityTypeList struct
type RoomAmenityTypeList struct {
	RoomAmenityType
	RoomAmenityCategoryNames *string `json:"room_amenity_category_names,omitempty"` // Room Amenity Category Names
}

// RoomAmenityTypeAPI Room Amenity Type API
type RoomAmenityTypeAPI struct {
	RoomAmenityTypeCode                *int                                    `json:"room_amenity_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"`                   // Room Amenity Type Code
	RoomAmenityTypeName                *string                                 `json:"room_amenity_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Adjoining rooms"` // Room Amenity Type Name
	RoomAmenityTypeAsset               *RoomAmenityTypeAssetAPI                `json:"room_amenity_type_asset,omitempty" gorm:"-"`                                                                                          // Room Amenity Type Asset
	RoomAmenityCategoryRoomAmenityType []RoomAmenityCategoryRoomAmenityTypeAPI `json:"room_amenity_category_room_amenity_type,omitempty" gorm:"-"`                                                                          // Room Amenity Category Room Amenity Type
}

// Seed RoomAmenityType
func (roomAmenityType *RoomAmenityType) Seed() *RoomAmenityType {
	seed := RoomAmenityType{
		RoomAmenityTypeAPI: RoomAmenityTypeAPI{
			RoomAmenityTypeCode: lib.Intptr(1),
			RoomAmenityTypeName: lib.Strptr("Adjoining rooms"),
			RoomAmenityTypeAsset: &RoomAmenityTypeAssetAPI{
				MultimediaDescriptionID: lib.UUIDPtr(uuid.New()),
			},
		},
	}
	_ = lib.Merge(seed, &roomAmenityType)
	return roomAmenityType
}
