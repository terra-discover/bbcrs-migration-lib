package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserTypeMenuLink Model
type UserTypeMenuLink struct {
	basic.Base
	basic.DataOwner
	UserTypeMenuLinkAPI
	MenuLink *MenuLink `json:"menu_link" gorm:"foreignKey:MenuLinkID;references:ID"`
}

// UserTypeMenuLinkAPI API
type UserTypeMenuLinkAPI struct {
	UserTypeID *uuid.UUID `json:"user_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	MenuLinkID *uuid.UUID `json:"menu_link_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsAllowed  *bool      `json:"is_allowed,omitempty"`
}
