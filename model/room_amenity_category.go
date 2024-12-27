package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// RoomAmenityCategory Amenity Category
type RoomAmenityCategory struct {
	basic.Base
	basic.DataOwner
	RoomAmenityCategoryAPI
	RoomAmenityCategoryTranslation *RoomAmenityCategoryTranslation `json:"amenity_category_translation,omitempty"`
	RoomAmenityCategoryAsset       *RoomAmenityCategoryAsset       `json:"room_amenity_category_asset,omitempty"` // Attraction Category Asset
}

// RoomAmenityCategoryAPI Amenity Category API
type RoomAmenityCategoryAPI struct {
	RoomAmenityCategoryName     *string                      `json:"room_amenity_category_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Bedroom"` // Room Amenity Category Name
	ParentRoomAmenityCategoryID *uuid.UUID                   `json:"parent_room_amenity_category_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                            // Parent Room Amenity Category ID
	IsDefault                   *bool                        `json:"is_default,omitempty" example:"true"`                                                                                             // Is Default
	Description                 *string                      `json:"description,omitempty" gorm:"type:text" example:"description"`                                                                    // Description
	RoomAmenityCategoryAsset    *RoomAmenityCategoryAssetAPI `json:"room_amenity_category_asset,omitempty" gorm:"-"`
}

// Seed data
func (a *RoomAmenityCategory) Seed() *[]RoomAmenityCategory {
	data := []RoomAmenityCategory{}
	items := []string{
		"Bedroom",
		"Bathroom",
		"Entertainment",
		"Food and drink",
	}

	for i := range items {
		var name string = items[i]
		data = append(data, RoomAmenityCategory{
			RoomAmenityCategoryAPI: RoomAmenityCategoryAPI{
				RoomAmenityCategoryName: &name,
				IsDefault:               lib.Boolptr(false),
				Description:             lib.Strptr(name),
			},
		})
	}

	return &data
}

// SeedOne AmenityCategory
func (a *RoomAmenityCategory) SeedOne() *RoomAmenityCategory {
	seed := RoomAmenityCategory{
		RoomAmenityCategoryAPI: RoomAmenityCategoryAPI{
			RoomAmenityCategoryName:     lib.Strptr("Bedroom"),
			ParentRoomAmenityCategoryID: lib.UUIDPtr(uuid.New()),
			RoomAmenityCategoryAsset: &RoomAmenityCategoryAssetAPI{
				MultimediaDescriptionID: lib.UUIDPtr(uuid.New()),
			},
		},
	}
	_ = lib.Merge(seed, &a)
	return a
}
