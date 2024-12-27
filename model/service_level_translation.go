package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceLevelTranslation Service Level Translation
type ServiceLevelTranslation struct {
	basic.Base
	basic.DataOwner
	ServiceLevelTranslationAPI
	ServiceLevelID *uuid.UUID `json:"service_level_id,omitempty" gorm:"type:varchar(36);uniqueIndex:service_level_translation_unique;not null" swaggertype:"string" format:"uuid"` // Service Level id
}

// ServiceLevelTranslationAPI Service Level Translation API
type ServiceLevelTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:service_level_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ServiceLevelName *string `json:"service_level_name,omitempty" gorm:"type:varchar(256)"`                                                             // Service Level Name
	Comment          *string `json:"comment,omitempty" gorm:"type:text"`                                                                                // Comment
	Description      *string `json:"description,omitempty" gorm:"type:text"`                                                                            // Description
}
