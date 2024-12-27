package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserTypeModuleCapability Model
type UserTypeModuleCapability struct {
	basic.Base
	basic.DataOwner
	UserTypeModuleCapabilityAPI
	ModuleCapability *ModuleCapability `json:"module_capability" gorm:"foreignKey:ModuleCapabilityID;references:ID"`
}

// UserTypeModuleCapabilityAPI API
type UserTypeModuleCapabilityAPI struct {
	UserTypeID         *uuid.UUID `json:"user_type_id,omitempty" gorm:"type:varchar(36);not null;index:idx_user_type_id,unique,where:deleted_at is null" format:"uuid"`
	ModuleCapabilityID *uuid.UUID `json:"module_capability_id,omitempty" gorm:"type:varchar(36);not null;index:idx_user_type_id,unique,where:deleted_at is null" format:"uuid"`
	IsAllowed          *bool      `json:"is_allowed,omitempty"`
}
