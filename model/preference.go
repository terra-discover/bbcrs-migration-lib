package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Preference Preference
type Preference struct {
	basic.Base
	basic.DataOwner
	PreferenceAPI
	PreferenceTranslation *PreferenceTranslation `json:"preference_translation,omitempty"`
	PreferenceType        *PreferenceType        `json:"preference_type" gorm:"foreignKey:PreferenceTypeID;references:ID"`
	FieldType             *FieldType             `json:"field_type" gorm:"foreignKey:FieldTypeID;references:ID"`
}

// PreferenceAPI Preference API
type PreferenceAPI struct {
	PreferenceName    *string    `json:"preference_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"queen bed"`                  // Preference Name
	PreferenceTypeID  *uuid.UUID `json:"preference_type_id,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Preference Type ID
	FieldTypeID       *uuid.UUID `json:"field_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                      // Field Type ID
	MaximumLength     *int       `json:"maximum_length,omitempty" example:"10"`                                                                                                   // Maximum Length
	DisplayLength     *int       `json:"display_length,omitempty" example:"10"`                                                                                                   // Display Length
	HideLabel         *bool      `json:"hide_label,omitempty" example:"true"`                                                                                                     // Hide Label
	MaskValue         *bool      `json:"mask_value,omitempty" example:"false"`                                                                                                    // Mask Value
	IsSameLine        *bool      `json:"is_same_line,omitempty" example:"true"`                                                                                                   // Is Same Line
	IsRequired        *bool      `json:"is_required,omitempty" example:"false"`                                                                                                   // Is Required
	IsEmailAddress    *bool      `json:"is_email_address,omitempty" example:"true"`                                                                                               // Is Email Address
	RegularExpression *string    `json:"regular_expression,omitempty" gorm:"type:text" example:"@[0-9]"`                                                                          // Regular Expression
	Description       *string    `json:"description,omitempty" gorm:"type:text" example:"deskripsi"`                                                                              // Description
}
