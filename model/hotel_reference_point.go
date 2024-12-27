package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelReferencePoint Hotel Reference Point
type HotelReferencePoint struct {
	basic.Base
	basic.DataOwner
	HotelReferencePointAPI
	ReferencePoint         *ReferencePoint         `json:"reference_point" gorm:"foreignKey:ReferencePointID;references:ID"` // ReferencePoint
	Hotel                  *Hotel                  `json:"hotel" gorm:"foreignKey:HotelID;references:ID"`                    // Hotel
	ReferencePointCategory *ReferencePointCategory `json:"reference_point_category" gorm:"foreignKey:ID"`
}

// HotelReferencePointAPI Hotel Reference Point API
type HotelReferencePointAPI struct {
	HotelID               *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`           // Hotel ID
	ReferencePointID      *uuid.UUID `json:"reference_point_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"` // Reference Point ID
	Direction             *string    `json:"direction,omitempty" gorm:"type:varchar(32)"`                                  // Direction
	Distance              *string    `json:"distance,omitempty" gorm:"type:varchar(32)"`                                   // Distance
	UnitOfMeasureID       *uuid.UUID `json:"unit_of_measure_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`           // Unit Of Measure ID
	IsNearest             *bool      `json:"is_nearest,omitempty"`                                                         // Is Nearest
	IsPrimary             *bool      `json:"is_primary,omitempty"`                                                         // Is Primary
	ToFrom                *string    `json:"to_from,omitempty" gorm:"type:varchar(16)"`                                    // To From
	IsApproximateDistance *bool      `json:"is_approximate_distance,omitempty"`                                            // Is Approximate Distance
}

// HotelReferencePointDataGet schema
type HotelReferencePointDataGet struct {
	ID                    *uuid.UUID       `json:"id,omitempty"`
	HotelID               *uuid.UUID       `json:"-"`
	ReferencePointID      *uuid.UUID       `json:"-"`
	CreatedAt             *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt             *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort                  *int64           `json:"sort,omitempty"`                                               // sort
	Status                *int64           `json:"status,omitempty"`
	Direction             *string          `json:"direction,omitempty"`               // Direction
	Distance              *string          `json:"distance,omitempty"`                // Distance
	UnitOfMeasureID       *uuid.UUID       `json:"unit_of_measure_id,omitempty"`      // Unit Of Measure ID
	IsNearest             *bool            `json:"is_nearest,omitempty"`              // Is Nearest
	IsPrimary             *bool            `json:"is_primary,omitempty"`              // Is Primary
	ToFrom                *string          `json:"to_from,omitempty"`                 // To From
	IsApproximateDistance *bool            `json:"is_approximate_distance,omitempty"` // Is Approximate Distance
	HotelReferencePointDataGetAPI
}

// HotelReferencePointDataGetAPI schema
type HotelReferencePointDataGetAPI struct {
	ReferencePoint         *ReferencePoint         `json:"reference_point,omitempty" gorm:"foreignKey:ReferencePointID;references:ID"`
	Hotel                  *Hotel                  `json:"hotel,omitempty" gorm:"foreignKey:HotelID;references:ID"`
	ReferencePointCategory *ReferencePointCategory `json:"reference_point_category" gorm:"foreignKey:ID"`
}
