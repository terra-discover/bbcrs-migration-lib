package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelRoomType Integration Partner Hotel Room Type
type IntegrationPartnerHotelRoomType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelID *uuid.UUID `json:"integration_partner_hotel_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_room_type__room_type_id,unique,where:deleted_at is null;not null"` // Integration Partner Hotel ID
	IntegrationPartnerHotelRoomTypeAPI
	IntegrationPartnerHotelRoomTypeTranslation *IntegrationPartnerHotelRoomTypeTranslation `json:"integration_partner_hotel_room_type_translation,omitempty"` // Integration Partner Hotel Room Type Translation
	IntegrationPartnerHotel                    *IntegrationPartnerHotel                    `json:"integration_partner_hotel,omitempty"`
	RoomType                                   *RoomType                                   `json:"room_type,omitempty"` // RoomType
}

// IntegrationPartnerHotelRoomTypeAPI Integration Partner Hotel Room Type API
type IntegrationPartnerHotelRoomTypeAPI struct {
	RoomTypeID   *uuid.UUID `json:"room_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_room_type__room_type_id,unique,where:deleted_at is null;not null"` // Room Type ID
	RoomTypeCode *string    `json:"room_type_code,omitempty" gorm:"type:varchar(256)"`                                                                                                                                      // Room Type Code
	RoomTypeName *string    `json:"room_type_name,omitempty" gorm:"type:varchar(512)"`                                                                                                                                      // Room Type Name
	MinOccupancy *int       `json:"min_occupancy,omitempty" gorm:"type:smallint"`                                                                                                                                           // Min Occupancy
	MaxOccupancy *int       `json:"max_occupancy,omitempty" gorm:"type:smallint"`                                                                                                                                           // Max Occupancy
	Description  *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                 // Description
	Comment      *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                     // Comment
}
