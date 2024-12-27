package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Form struct Form
type Form struct {
	basic.Base
	basic.DataOwner
	FormAPI
	FormType        *FormType        `json:"form_type,omitempty"`
	FormTranslation *FormTranslation `json:"form_translation,omitempty"`
}

// FormAPI Form API
type FormAPI struct {
	FormCode    *string    `json:"form_code,omitempty" gorm:"type:varchar(36);not null;index:idx_form_code_deleted_at,unique,where:deleted_at is null" example:"contact-us"`  // Form Code
	FormName    *string    `json:"form_name,omitempty" gorm:"type:varchar(256);not null;index:idx_form_name_deleted_at,unique,where:deleted_at is null" example:"Contact Us"` // Form Name
	FormTypeID  *uuid.UUID `json:"form_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                         // Form Type ID
	Opening     *string    `json:"opening,omitempty" gorm:"type:text"`                                                                                                        // Opening
	Closing     *string    `json:"closing,omitempty" gorm:"type:text"`                                                                                                        // Closing
	Excerpt     *string    `json:"excerpt,omitempty" gorm:"type:text"`                                                                                                        // Excerpt
	Description *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                    // Description
}

// Seed Form
func (form *Form) Seed() *Form {
	seed := Form{
		FormAPI: FormAPI{
			FormCode:    lib.Strptr("contact-us"),
			FormName:    lib.Strptr("Contact Us"),
			FormTypeID:  lib.UUIDPtr(uuid.New()),
			Opening:     nil,
			Closing:     nil,
			Excerpt:     nil,
			Description: nil,
		},
	}
	_ = lib.Merge(seed, &form)
	return form
}
