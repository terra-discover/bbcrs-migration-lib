package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PreferenceType Preference Type
type PreferenceType struct {
	basic.Base
	basic.DataOwner
	PreferenceTypeAPI
	PreferenceTypeTranslation *PreferenceTypeTranslation `json:"preference_type_translation,omitempty"` // Preference Type Translation
}

// PreferenceTypeAPI Preference Type API
type PreferenceTypeAPI struct {
	PreferenceTypeName     *string    `json:"preference_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // Preference Type Name
	ParentPreferenceTypeID *uuid.UUID `json:"parent_preference_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // Parent Preference Type ID
	Description            *string    `json:"description,omitempty" gorm:"type:text"`                                                                  // Description
}
