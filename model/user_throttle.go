package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserThrottle User Throttle
type UserThrottle struct {
	basic.Base
	basic.DataOwner
	UserThrottleAPI
	UserAccount *UserAccount `json:"user_account" gorm:"foreignKey:UserAccountID;references:ID"`
}

// UserThrottleAPI User Throttle Api
type UserThrottleAPI struct {
	UserAccountID *uuid.UUID       `json:"user_account_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	IPAddress     *string          `json:"ip_address,omitempty" gorm:"type:varchar(256)"`
	LoginAttempts *int             `json:"login_attempts,omitempty" gorm:"type:int"`
	LastAttemptAt *strfmt.DateTime `json:"last_attempt_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	Requests      *int             `json:"requests,omitempty" gorm:"type:int"`
	RequestUntil  *strfmt.DateTime `json:"request_until,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	IsSuspended   *bool            `json:"is_suspended,omitempty"`
	SuspendedAt   *strfmt.DateTime `json:"suspended_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	IsBanned      *bool            `json:"is_banned,omitempty"`
	BannedAt      *strfmt.DateTime `json:"banned_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
}
