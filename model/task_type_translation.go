package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TaskTypeTranslation Task Type Translation
type TaskTypeTranslation struct {
	basic.Base
	basic.DataOwner
	TaskTypeTranslationAPI
	TaskTypeID *uuid.UUID `json:"task_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:task_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Task Type id
}

// TaskTypeTranslationAPI Task Type Translation API
type TaskTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:task_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	TaskTypeName *string `json:"task_type_name,omitempty" gorm:"type:varchar(256)" example:"General"`                                           // Task Type Name
}
