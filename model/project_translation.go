package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProjectTranslation Project Translation
type ProjectTranslation struct {
	basic.Base
	basic.DataOwner
	ProjectTranslationAPI
	ProjectID *uuid.UUID `json:"project_id,omitempty" gorm:"type:varchar(36);uniqueIndex:project_translation_unique;not null" swaggertype:"string" format:"uuid"` // Project id
}

// ProjectTranslationAPI Project Translation API
type ProjectTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:project_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ProjectName  *string `json:"project_name,omitempty" gorm:"type:varchar(256)"`                                                             // Project Name
}
