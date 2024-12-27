package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FormTypeTranslation Form Type Translation
type FormTypeTranslation struct {
	basic.Base
	basic.DataOwner
	FormTypeTranslationAPI
	FormTypeID *uuid.UUID `json:"form_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:form_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Form Type id
}

// FormTypeTranslationAPI Form Type Translation API
type FormTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:form_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FormTypeName *string `json:"form_type_name,omitempty" gorm:"type:varchar(256)" example:"Survey"`                                            // Form Type Name
}
