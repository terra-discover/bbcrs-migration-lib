package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelRoomTypeOccupancy Integration Partner HotelRoomTypeOccupancy
type IntegrationPartnerHotelRoomTypeOccupancy struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelRoomTypeOccupancyAPI
	IntegrationPartnerHotel *IntegrationPartnerHotel `json:"integration_partner_hotel,omitempty"`
	RoomType                *RoomType                `json:"room_type,omitempty"`
	AgeQualifyingType       *AgeQualifyingType       `json:"age_qualifying_type,omitempty"`
}

// IntegrationPartnerHotelRoomTypeOccupancyAPI Integration PartnerHotelRoomTypeOccupancy API
type IntegrationPartnerHotelRoomTypeOccupancyAPI struct {
	IntegrationPartnerHotelID *uuid.UUID `json:"integration_partner_hotel_id,omitempty" gorm:"type:varchar(36);not null;"`
	RoomTypeID                *uuid.UUID `json:"room_type_id,omitempty" gorm:"type:varchar(36);not null;"`
	AgeQualifyingTypeID       *uuid.UUID `json:"age_qualifying_type_id,omitempty" gorm:"type:varchar(36);not null;"`
	MinOccupancy              *int       `json:"min_occupancy,omitempty" gorm:"type:smallint"`
	MaxOccupancy              *int       `json:"max_occupancy,omitempty" gorm:"type:smallint"`
	Comment                   *string    `json:"comment,omitempty" gorm:"type:text"`
}
