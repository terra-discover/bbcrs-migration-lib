package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelAttraction Agent Hotels Attractions
type HotelAttraction struct {
	basic.Base
	basic.DataOwner
	HotelAttractionAPI
	Attraction         *Attraction         `json:"attraction" gorm:"foreignKey:AttractionID;references:ID"` // Attraction
	Hotel              *Hotel              `json:"hotel" gorm:"foreignKey:HotelID;references:ID"`           // Hotel
	AttractionCategory *AttractionCategory `json:"attraction_category" gorm:"foreignKey:ID"`
}

// HotelAttractionAPI Agent Hotels Attractions API
type HotelAttractionAPI struct {
	HotelID               *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`      // Hotel ID
	AttractionID          *uuid.UUID `json:"attraction_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"` // Attraction ID
	ProximityID           *uuid.UUID `json:"proximity_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`            // Proximity ID
	OperationScheduleID   *uuid.UUID `json:"operation_schedule_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`   // Operation Schedule ID
	Direction             *string    `json:"direction,omitempty" gorm:"type:varchar(32)"`                             // Direction
	Distance              *string    `json:"distance,omitempty" gorm:"type:varchar(32)"`                              // Distance
	UnitOfMeasureID       *uuid.UUID `json:"unit_of_measure_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`      // Unit Of Measure ID
	IsNearest             *bool      `json:"is_nearest,omitempty"`                                                    // Is Nearest
	IsPrimary             *bool      `json:"is_primary,omitempty"`                                                    // Is Primary
	ToFrom                *string    `json:"to_from,omitempty" gorm:"type:varchar(16)"`                               // To From
	IsApproximateDistance *bool      `json:"is_approximate_distance,omitempty"`                                       // Is Approximate Distance
}

// HotelAttractionDataGet schema
type HotelAttractionDataGet struct {
	ID                    *uuid.UUID       `json:"id,omitempty"`
	HotelID               *uuid.UUID       `json:"-"`
	AttractionID          *uuid.UUID       `json:"-"`
	CreatedAt             *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt             *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort                  *int64           `json:"sort,omitempty"`                                               // sort
	Status                *int64           `json:"status,omitempty"`
	ProximityID           *uuid.UUID       `json:"proximity_id,omitempty"`            // Proximity ID
	OperationScheduleID   *uuid.UUID       `json:"operation_schedule_id,omitempty"`   // Operation Schedule ID
	Direction             *string          `json:"direction,omitempty"`               // Direction
	Distance              *string          `json:"distance,omitempty"`                // Distance
	UnitOfMeasureID       *uuid.UUID       `json:"unit_of_measure_id,omitempty"`      // Unit Of Measure ID
	IsNearest             *bool            `json:"is_nearest,omitempty"`              // Is Nearest
	IsPrimary             *bool            `json:"is_primary,omitempty"`              // Is Primary
	ToFrom                *string          `json:"to_from,omitempty"`                 // To From
	IsApproximateDistance *bool            `json:"is_approximate_distance,omitempty"` // Is Approximate Distance
	HotelAttractionDataGetAPI
}

// HotelAttractionDataGetAPI schema
type HotelAttractionDataGetAPI struct {
	Attraction         *Attraction         `json:"attraction,omitempty" gorm:"foreignKey:AttractionID;references:ID"`
	Hotel              *Hotel              `json:"hotel,omitempty" gorm:"foreignKey:HotelID;references:ID"`
	AttractionCategory *AttractionCategory `json:"attraction_category" gorm:"foreignKey:ID"`
}
