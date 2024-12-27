package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PreferenceOption Preference Type
type PreferenceOption struct {
	basic.Base
	basic.DataOwner
	PreferenceOptionAPI
	Preference                  *Preference                  `json:"preference" gorm:"foreignKey:PreferenceID;references:ID"`
	PreferenceOptionTranslation *PreferenceOptionTranslation `json:"preference_option_translation,omitempty"` // Preference Option Translation
}

// PreferenceOptionAPI Preference Type API
type PreferenceOptionAPI struct {
	PreferenceOptionName *string    `json:"preference_option_name,omitempty" gorm:"type:varchar(256);not null"`                          // Preference Option Name
	PreferenceID         *uuid.UUID `json:"preference_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Preference ID
	IsDefault            *bool      `json:"is_default,omitempty" example:"true"`
	Description          *string    `json:"description,omitempty" gorm:"type:text"` // Description
}
