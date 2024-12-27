package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// HotelRating model
type HotelRating struct {
	basic.Base
	basic.DataOwner
	HotelRatingAPI
	RatingType *RatingType `json:"rating_type,omitempty"`
}

// HotelRatingAPI struct
type HotelRatingAPI struct {
	HotelID      *uuid.UUID `json:"hotel_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`       // Hotel ID
	RatingTypeID *uuid.UUID `json:"rating_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // Rating Type ID
	Rating       *float64   `json:"rating,omitempty"`                                                                    // Rating
}
