package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EnrollmentMethodTranslation Enrollment Method Translation
type EnrollmentMethodTranslation struct {
	basic.Base
	basic.DataOwner
	EnrollmentMethodTranslationAPI
	EnrollmentMethodID *uuid.UUID `json:"enrollment_method_id,omitempty" gorm:"type:varchar(36);uniqueIndex:enrollment_method_translation_unique;not null" swaggertype:"string" format:"uuid"` // Enrollment Method id
}

// EnrollmentMethodTranslationAPI Enrollment Method Translation API
type EnrollmentMethodTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:enrollment_method_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	EnrollmentMethodName *string `json:"enrollment_method_name,omitempty" gorm:"type:varchar(256)" example:"Online"`                                            // Enrollment Method Name
}
