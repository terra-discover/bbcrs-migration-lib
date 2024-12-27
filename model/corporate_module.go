package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type CorporateModule struct {
	basic.Base
	basic.DataOwner
	CorporateModuleAPI
}

type CorporateModuleAPI struct {
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);index:,where:deleted_at is null;not null;" swaggertype:"string" format:"uuid"`
	ModuleID    *uuid.UUID `json:"module_id,omitempty" gorm:"type:varchar(36);index:,where:deleted_at is null;not null;" swaggertype:"string" format:"uuid"`
}
