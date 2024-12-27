package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FormFieldOptionTranslation FormFieldOption Translation
type FormFieldOptionTranslation struct {
	basic.Base
	basic.DataOwner
	FormFieldOptionTranslationAPI
	FormFieldOptionID *uuid.UUID `json:"form_field_option_id,omitempty" gorm:"type:varchar(36);uniqueIndex:form_field_option_translation_unique;not null" swaggertype:"string" format:"uuid"` // FormFieldOption id
}

// FormFieldOptionTranslationAPI FormFieldOption Translation API
type FormFieldOptionTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:form_field_option_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FormFieldOptionName *string `json:"form_field_option_name,omitempty" gorm:"type:varchar(256)"`
	Description         *string `json:"description,omitempty" gorm:"type:text"` // FormFieldOption Name
}
