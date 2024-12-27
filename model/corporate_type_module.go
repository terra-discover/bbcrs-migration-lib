package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type CorporateTypeModule struct {
	basic.Base
	basic.DataOwner
	CorporateTypeAPI
}

type CorporateTypeModuleAPI struct {
	CorporateTypeID *uuid.UUID `json:"corporate_type_id,omitempty" gorm:"type:varchar(36);index:,where:deleted_at is null;not null;" swaggertype:"string" format:"uuid"`
	ModuleID        *uuid.UUID `json:"module_id,omitempty" gorm:"type:varchar(36);index:,where:deleted_at is null;not null;" swaggertype:"string" format:"uuid"`
}
