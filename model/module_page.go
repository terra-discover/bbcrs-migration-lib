package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ModulePage Module read only model
type ModulePage struct {
	basic.Base
	basic.DataOwner
	ModulePageAPI
	Module *Module `json:"module,omitempty" gorm:"foreignKey:ModuleID;references:ID"`
	Page   *Page   `json:"page,omitempty" gorm:"foreignKey:PageID;references:ID"`
}

// ModulePageAPI Module API
type ModulePageAPI struct {
	ModuleID *uuid.UUID `json:"module_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
	PageID   *uuid.UUID `json:"page_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);"`
}
