package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserSocialNetwork Model
type UserSocialNetwork struct {
	basic.Base
	basic.DataOwner
	UserSocialNetworkAPI
	SocialNetwork *SocialNetwork `json:"social_network" gorm:"foreignKey:SocialNetworkID;references:ID"`
	UserAccount   *UserAccount   `json:"user_account" gorm:"foreignKey:UserAccountID;references:ID"`
}

// UserSocialNetworkAPI API
type UserSocialNetworkAPI struct {
	UserAccountID   *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	SocialNetworkID *uuid.UUID `json:"social_network_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
