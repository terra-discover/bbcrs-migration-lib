package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PersonFormField PersonFormField
type PersonFormField struct {
	basic.Base
	basic.DataOwner
	PersonFormFieldAPI
	Person    *Person    `json:"person" gorm:"foreignKey:PersonID;references:ID"`
	FormField *FormField `json:"form_field" gorm:"foreignKey:FormFieldID;references:ID"`
}

// PersonFormFieldAPI PersonFormField API
type PersonFormFieldAPI struct {
	PersonID              *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	FormFieldID           *uuid.UUID `json:"form_field_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	FormFieldValue        *string    `json:"form_field_value,omitempty" gorm:"type:text"`
	HasFormFieldSelection *bool      `json:"has_form_field_selection,omitempty" example:"true"`
}
