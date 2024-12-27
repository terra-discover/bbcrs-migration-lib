package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserTypeUser Model
type UserTypeUser struct {
	basic.Base
	basic.DataOwner
	UserTypeUserAPI
	UserAccount *UserAccount `json:"user_account" gorm:"foreignKey:UserAccountID;references:ID"`
}

// UserTypeUserAPI API
type UserTypeUserAPI struct {
	UserAccountID *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	UserTypeID    *uuid.UUID `json:"user_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
