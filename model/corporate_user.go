package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateUser Model
type CorporateUser struct {
	basic.Base
	basic.DataOwner
	CorporateUserAPI
	Corporate   *Corporate   `json:"corporate,omitempty"`
	UserAccount *UserAccount `json:"user_account,omitempty"`
}

// CorporateUserAPI Model
type CorporateUserAPI struct {
	CorporateID   *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`    // Corporate ID
	UserAccountID *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // User Account ID
	IsPrimary     *bool      `json:"is_primary,omitempty"`
}
