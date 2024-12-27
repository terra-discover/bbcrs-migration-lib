package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PersonPreferenceSelection PersonPreferenceSelection
type PersonPreferenceSelection struct {
	basic.Base
	basic.DataOwner
	PersonPreferenceSelectionAPI
	PersonPreference *PersonPreference `json:"person_preference" gorm:"foreignKey:PersonPreferenceID;references:ID"`
}

// PersonPreferenceSelectionAPI PersonPreferenceSelection API
type PersonPreferenceSelectionAPI struct {
	PersonPreferenceID       *uuid.UUID `json:"person_preference_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	PreferenceOptionID       *uuid.UUID `json:"preference_option_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	PreferenceSelectionValue *string    `json:"preference_selection_value,omitempty" gorm:"type:text"`
}
