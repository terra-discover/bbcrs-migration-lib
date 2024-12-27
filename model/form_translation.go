package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FormTranslation Form Translation
type FormTranslation struct {
	basic.Base
	basic.DataOwner
	FormTranslationAPI
	FormID *uuid.UUID `json:"form_id,omitempty" gorm:"type:varchar(36);uniqueIndex:form_translation_unique;not null" swaggertype:"string" format:"uuid"` // Form id
}

// FormTranslationAPI Form Translation API
type FormTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:form_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	FormName     *string `json:"form_name,omitempty" gorm:"type:varchar(256)" example:"Contact Us"`                                        // Form Name
	Description  *string `json:"description,omitempty"`                                                                                    // Description
}
