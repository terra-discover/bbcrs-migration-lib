package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AttractionCategoryAttraction struct
type AttractionCategoryAttraction struct {
	basic.Base
	basic.DataOwner
	AttractionCategoryAttractionAPI
	Description        *string             `json:"description,omitempty" gorm:"type:text"` // Description
	Attraction         *Attraction         `json:"attraction,omitempty"`
	AttractionCategory *AttractionCategory `json:"attraction_category,omitempty"` // Hotel Amenity Category
}

// AttractionCategoryAttractionAPI for handle request body
type AttractionCategoryAttractionAPI struct {
	AttractionID         *uuid.UUID `json:"attraction_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`          // Hotel Amenity Type ID
	AttractionCategoryID *uuid.UUID `json:"attraction_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Hotel Amenity Category ID
}
