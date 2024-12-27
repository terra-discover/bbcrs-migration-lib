package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// NamePrefix Name Prefix
type NamePrefix struct {
	basic.Base
	basic.DataOwner
	NamePrefixAPI
	NamePrefixTranslation *NamePrefixTranslation `json:"name_prefix_translation,omitempty"`                         // Name Prefix Translation
	Gender                *Gender                `json:"gender,omitempty" gorm:"foreignKey:GenderID;references:ID"` // NamePrefix Gender
}

// NamePrefixAPI Name Prefix API
type NamePrefixAPI struct {
	NamePrefixCode *string    `json:"name_prefix_code,omitempty" gorm:"type:varchar(16);not null;index:,unique,where:deleted_at is null"`  // Name Prefix Code
	NamePrefixName *string    `json:"name_prefix_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Name Prefix Name
	GenderID       *uuid.UUID `json:"gender_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                      // Gender ID
}
