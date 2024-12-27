package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PersonFormFieldSelection PersonFormFieldSelection
type PersonFormFieldSelection struct {
	basic.Base
	basic.DataOwner
	PersonFormFieldSelectionAPI
	PersonFormField *PersonFormField `json:"person_form_field" gorm:"foreignKey:PersonFormFieldID;references:ID"`
	FormFieldOption *FormFieldOption `json:"form_field_option" gorm:"foreignKey:FormFieldOptionID;references:ID"`
}

// PersonFormFieldSelectionAPI PersonFormFieldSelection API
type PersonFormFieldSelectionAPI struct {
	PersonFormFieldID       *uuid.UUID `json:"person_form_field_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	FormFieldOptionID       *uuid.UUID `json:"form_field_option_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	FormFieldSelectionValue *string    `json:"form_field_selection_value,omitempty" gorm:"type:text"`
}
