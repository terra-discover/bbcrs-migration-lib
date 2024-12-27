package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FormField FormField
type FormField struct {
	basic.Base
	basic.DataOwner
	FormFieldAPI
	FormFieldTranslation *FormFieldTranslation `json:"form_field_translation,omitempty"`
	FieldType            *FieldType            `json:"field_type" gorm:"foreignKey:FieldTypeID;references:ID"`
}

// FormFieldAPI FormField API
type FormFieldAPI struct {
	FormFieldName     *string    `json:"form_field_name,omitempty" gorm:"type:varchar(256);not null"`
	FormID            *uuid.UUID `json:"form_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	FieldTypeID       *uuid.UUID `json:"field_type_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	MaximumLength     *int       `json:"maximum_length,omitempty" gorm:"type:smallint" example:"1"`
	DisplayLength     *int       `json:"display_length,omitempty" gorm:"type:smallint" example:"1"`
	HideLabel         *bool      `json:"hide_label,omitempty"  example:"true"`
	MaskValue         *bool      `json:"mask_value,omitempty" example:"true"`
	IsSameLine        *bool      `json:"is_same_line,omitempty" example:"true"`
	IsRequired        *bool      `json:"is_required,omitempty" example:"true"`
	IsEmailAddress    *bool      `json:"is_email_address,omitempty" example:"true"`
	RegularExpression *string    `json:"regular_expression,omitempty" gorm:"type:text"`
	Description       *string    `json:"description,omitempty" gorm:"type:text"`
}
