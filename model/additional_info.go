package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdditionalInfo Additional Info
type AdditionalInfo struct {
	basic.Base
	basic.DataOwner
	AdditionalInfoAPI
	AdditionalInfoTranslation *AdditionalInfoTranslation `json:"additional_info_translation,omitempty"` // Additional Info Translation
	FieldType                 *FieldType                 `json:"field_type,omitempty"`                  // Field Type
	AdditionalInfoOption      []AdditionalInfoOption     `json:"additional_info_option,omitempty"`
}

// AdditionalInfoAPI Additional Info API
type AdditionalInfoAPI struct {
	AdditionalInfoName *string    `json:"additional_info_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Additional Info Name
	FieldTypeID        *uuid.UUID `json:"field_type_id,omitempty" swaggertype:"string" format:"uuid"`                                              // Field Type ID
	MaximumLength      *int       `json:"maximum_length,omitempty" gorm:"type:smallint"`                                                           // Maximum Length
	DisplayLength      *int       `json:"display_length,omitempty" gorm:"type:smallint"`                                                           // Display Length
	HideLabel          *bool      `json:"hide_label,omitempty"`                                                                                    // Hide Label
	MaskValue          *bool      `json:"mask_value,omitempty"`                                                                                    // Mask Value
	IsSameLine         *bool      `json:"is_same_line,omitempty"`                                                                                  // Is Same Line
	IsRequired         *bool      `json:"is_required,omitempty"`                                                                                   // Is Required
	IsEmailAddress     *bool      `json:"is_email_address,omitempty"`                                                                              // Is Email Address
	RegularExpression  *string    `json:"regular_expression,omitempty" gorm:"type:text"`                                                           // Regular Expression
	Description        *string    `json:"description,omitempty" gorm:"type:text"`                                                                  // Description
}
