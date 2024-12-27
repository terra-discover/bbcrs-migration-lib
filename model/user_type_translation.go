package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// UserTypeTranslation User Type Translation
type UserTypeTranslation struct {
	basic.Base
	basic.DataOwner
	UserTypeTranslationAPI
	UserTypeID *uuid.UUID `json:"user_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:user_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // User Type id
}

// UserTypeTranslationAPI User Type Translation API
type UserTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:user_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	UserTypeName *string `json:"user_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // User Type Name                                                                       // Assign All Menu Links
	Comment      *string `json:"comment,omitempty"`                                                                                             // Comment
	Description  *string `json:"description,omitempty"`                                                                                         // Description
}
