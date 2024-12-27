package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ModuleCapability Model
type ModuleCapability struct {
	basic.Base
	basic.DataOwner
	ModuleCapabilityAPI
	Capability *Capability `json:"capability" gorm:"foreignKey:CapabilityID;references:ID"`
	Module     *Module     `json:"module" gorm:"foreignKey:ModuleID;references:ID"`
}

// ModuleCapabilityAPI API
type ModuleCapabilityAPI struct {
	ModuleID     *uuid.UUID `json:"module_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	CapabilityID *uuid.UUID `json:"capability_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
