package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserMenuLink Model
type UserMenuLink struct {
	basic.Base
	basic.DataOwner
	UserMenuLinkAPI
	MenuLink    *MenuLink    `json:"menu_link" gorm:"foreignKey:MenuLinkID;references:ID"`
	UserAccount *UserAccount `json:"user_account" gorm:"foreignKey:UserAccountID;references:ID"`
}

// UserMenuLinkAPI API
type UserMenuLinkAPI struct {
	UserAccountID *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	MenuLinkID    *uuid.UUID `json:"menu_link_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsAllowed     *bool      `json:"is_allowed,omitempty"`
}
