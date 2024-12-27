package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CapabilityCategory Capability Category
type CapabilityCategory struct {
	basic.Base
	basic.DataOwner
	CapabilityCategoryAPI
	CapabilityCategoryTranslation *CapabilityCategoryTranslation `json:"capability_category_translation,omitempty"` // Capability Category Translation
}

// CapabilityCategoryAPI Capability Category API
type CapabilityCategoryAPI struct {
	CapabilityCategoryCode     *string    `json:"capability_category_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`  // Capability Category Code
	CapabilityCategoryName     *string    `json:"capability_category_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Capability Category Name
	ParentCapabilityCategoryID *uuid.UUID `json:"parent_capability_category_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // Parent Capability Category ID
	IsDefault                  *bool      `json:"is_default,omitempty"`                                                                                        // Is Default
	Description                *string    `json:"description,omitempty" gorm:"type:text"`                                                                      // Description
}
