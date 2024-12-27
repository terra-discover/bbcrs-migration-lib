package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomType Model
type RoomType struct {
	basic.Base
	basic.DataOwner
	RoomTypeAPI
	Hotel               *Hotel               `json:"hotel" gorm:"foreignKey:HotelID;references:ID"`
	SegmentCategory     *SegmentCategory     `json:"segment_category" gorm:"foreignKey:SegmentCategoryID;references:ID"`
	RoomLocationType    *RoomLocationType    `json:"room_location_type" gorm:"foreignKey:RoomLocationTypeID;references:ID"`
	RoomViewType        *RoomViewType        `json:"room_view_type" gorm:"foreignKey:RoomViewTypeID;references:ID"`
	RoomClassification  *RoomClassification  `json:"room_classification" gorm:"foreignKey:RoomClassificationID;references:ID"`
	RoomTypeTranslation *RoomTypeTranslation `json:"room_type_translation,omitempty"`
}

// RoomTypeAPI API
type RoomTypeAPI struct {
	HotelID                 *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	RoomTypeCode            *string    `json:"room_type_code,omitempty" gorm:"type:varchar(36);not null"`
	RoomTypeName            *string    `json:"room_type_name,omitempty" gorm:"type:varchar(256);not null"`
	SegmentCategoryID       *uuid.UUID `json:"segment_category_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	RoomLocationTypeID      *uuid.UUID `json:"room_location_type_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	RoomViewTypeID          *uuid.UUID `json:"room_view_type_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	RoomClassificationID    *uuid.UUID `json:"room_classification_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	NumberOfBeds            *int64     `json:"number_of_beds" example:"10"`
	IsNonSmoking            *bool      `json:"is_non_smoking,omitempty"`
	RoomSize                *float64   `json:"room_size,omitempty" gorm:"type:decimal(19,4)" example:"25.00"`
	RoomSizeUnitOfMeasureID *uuid.UUID `json:"room_size_unit_of_measure_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	NumberOfRooms           *int64     `json:"number_of_rooms" example:"10"`
	MaxOccupancy            *int64     `json:"max_occupancy" example:"10"`
	IsAccrual               *bool      `json:"is_accrual,omitempty"`
	IsRedeemable            *bool      `json:"is_redeemable,omitempty"`
	Description             *string    `json:"description,omitempty" gorm:"type:text"`
}
