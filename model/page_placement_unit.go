package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PagePlacementUnit struct
type PagePlacementUnit struct {
	basic.Base
	basic.DataOwner
	PagePlacementUnitAPI
	Page          *Page          `json:"page,omitempty"`           // Page
	BuiltInType   *BuiltInType   `json:"built_in_type,omitempty"`  // Built In Type
	ThemeLocation *ThemeLocation `json:"theme_location,omitempty"` // Theme Location
}

// PagePlacementUnitAPI struct
type PagePlacementUnitAPI struct {
	PageID          *uuid.UUID `json:"page_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`           // Page ID
	BuiltInTypeID   *uuid.UUID `json:"built_in_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`  // Built Type ID
	ThemeLocationID *uuid.UUID `json:"theme_location_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Theme Location ID
}
