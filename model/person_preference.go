package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PersonPreference PersonPreference
type PersonPreference struct {
	basic.Base
	basic.DataOwner
	PersonPreferenceAPI
	Person *Person `json:"person" gorm:"foreignKey:PersonID;references:ID"`
}

// PersonPreferenceAPI PersonPreference API
type PersonPreferenceAPI struct {
	PersonID               *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	PreferenceID           *uuid.UUID `json:"preference_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	PreferenceValue        *string    `json:"preference_value,omitempty" gorm:"type:text"`
	HasPreferenceSelection *bool      `json:"has_preference_selection,omitempty" example:"true"`
}
