package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateMenu model
type CorporateMenu struct {
	basic.Base
	basic.DataOwner
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_menu_id,unique,where:deleted_at is null"` // Corporate Id
	MenuID      *uuid.UUID `json:"menu_id,omitempty" gorm:"type:varchar(36);index:idx_corporate_menu_id,unique,where:deleted_at is null"`      // Menu Id
}
