package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// JobTitleTranslation JobTitle Translation
type JobTitleTranslation struct {
	basic.Base
	basic.DataOwner
	JobTitleTranslationAPI
	JobTitleID *uuid.UUID `json:"job_title_id,omitempty" gorm:"type:varchar(36);uniqueIndex:job_title_translation_unique;not null" swaggertype:"string" format:"uuid"` // Job Title id
}

// JobTitleTranslationAPI JobTitle Translation API
type JobTitleTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:job_title_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	JobTitleName *string `json:"job_title_name,omitempty" gorm:"type:varchar(256)"`                                                             // Job Title Name
}
