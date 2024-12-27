package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserSetting Model
type UserSetting struct {
	basic.Base
	basic.DataOwner
	UserSettingAPI
	UserAccount *UserAccount `json:"-" gorm:"foreignKey:UserAccountID;references:ID" swaggerignore:"true"`
}

// UserSettingAPI API
type UserSettingAPI struct {
	UserAccountID             *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ReceiveSystemNotification *bool      `json:"receive_system_notification,omitempty"`
	ReceiveTravelDeals        *bool      `json:"receive_travel_deals,omitempty"`
	ReceiveOtherInformation   *bool      `json:"receive_other_information,omitempty"`
}
