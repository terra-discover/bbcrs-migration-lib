package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserModuleCapability Model
type UserModuleCapability struct {
	basic.Base
	basic.DataOwner
	UserModuleCapabilityAPI
	ModuleCapability *ModuleCapability `json:"module_capability" gorm:"foreignKey:ModuleCapabilityID;references:ID"`
	UserAccount      *UserAccount      `json:"user_account" gorm:"foreignKey:UserAccountID;references:ID"`
}

// UserModuleCapabilityAPI API
type UserModuleCapabilityAPI struct {
	UserAccountID      *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ModuleCapabilityID *uuid.UUID `json:"module_capability_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsAllowed          *bool      `json:"is_allowed,omitempty"`
}
