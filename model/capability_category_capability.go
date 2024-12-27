package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CapabilityCategoryCapability Capability Category Capability
type CapabilityCategoryCapability struct {
	basic.Base
	basic.DataOwner
	CapabilityCategoryCapabilityAPI
}

// CapabilityCategoryCapabilityAPI Capability Category Capability Api
type CapabilityCategoryCapabilityAPI struct {
	CapabilityCategoryID *uuid.UUID `json:"capability_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Capability Category Id
	CapabilityID         *uuid.UUID `json:"capability_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`          // Capability Id
}
