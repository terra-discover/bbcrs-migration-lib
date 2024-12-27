package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FormFieldOption FormFieldOption
type FormFieldOption struct {
	basic.Base
	basic.DataOwner
	FormFieldOptionAPI
	FormFieldOptionTranslation *FormFieldOptionTranslation `json:"form_field_option_translation,omitempty"`
	FormField                  *FormField                  `json:"form_field" gorm:"foreignKey:FormFieldID;references:ID"`
}

// FormFieldOptionAPI FormFieldOption API
type FormFieldOptionAPI struct {
	FormFieldOptionName *string    `json:"form_field_option_name,omitempty" gorm:"type:varchar(256);not null"`
	FormFieldID         *uuid.UUID `json:"form_field_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`
	IsDefault           *bool      `json:"is_default,omitempty"  example:"true"`
	Description         *string    `json:"description,omitempty" gorm:"type:text"`
}
