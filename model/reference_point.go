package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReferencePoint Reference Point
type ReferencePoint struct {
	basic.Base
	basic.DataOwner
	ReferencePointAPI
	ReferencePointTranslation *ReferencePointTranslation `json:"reference_point_translation,omitempty"` // Reference Point Translation
	ReferencePointCategory    *ReferencePointCategory    `json:"reference_point_category,omitempty"`    // Reference Point Category
}

// ReferencePointAPI Reference Point API
type ReferencePointAPI struct {
	ReferencePointName       *string    `json:"reference_point_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Reference Point Name
	BusinessEntityID         *uuid.UUID `json:"business_entity_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                 // Business Entity ID
	ReferencePointCategoryID *uuid.UUID `json:"reference_point_category_id,omitempty" swaggertype:"string" format:"uuid"`                                // Reference Point Category ID
	Description              *string    `json:"description,omitempty" gorm:"type:text"`                                                                  // Description
}
