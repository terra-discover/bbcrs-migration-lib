package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FormFieldTranslation FormField Translation
type FormFieldTranslation struct {
	basic.Base
	basic.DataOwner
	FormFieldTranslationAPI
	FormFieldID *uuid.UUID `json:"form_field_id,omitempty" gorm:"type:varchar(36);uniqueIndex:form_field_translation_unique;not null" swaggertype:"string" format:"uuid"` // FormField id
}

// FormFieldTranslationAPI FormField Translation API
type FormFieldTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:form_field_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FormFieldName *string `json:"form_field_name,omitempty" gorm:"type:varchar(256)"`
	Description   *string `json:"description,omitempty" gorm:"type:text"` // FormField Name
}
