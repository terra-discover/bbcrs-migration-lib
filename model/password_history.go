package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PasswordHistory struct
type PasswordHistory struct {
	basic.Base
	basic.DataOwner
	PasswordHistoryAPI
}

// PasswordHistoryAPI struct
type PasswordHistoryAPI struct {
	UserAccountID     *uuid.UUID       `json:"person_id,omitempty" gorm:"type:varchar(36);not null;"`                                 // UserAccountID
	IPAddress         *string          `json:"ip_address,omitempty" gorm:"type:varchar(256)"`                                         // Ip Address
	Password          *string          `json:"password,omitempty" gorm:"type:varchar(256);not null;"`                                 // Password
	Salt              *string          `json:"salt,omitempty" gorm:"type:varchar(256);not null;"`                                     // Salt
	IsSystemGenerated *bool            `json:"is_system_generated,omitempty" gorm:"default:false"`                                    // Is System Generated
	LastChange        *strfmt.DateTime `json:"last_change,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // Last Change
	Expiration        *strfmt.DateTime `json:"expiration,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`  // Expiration
}
