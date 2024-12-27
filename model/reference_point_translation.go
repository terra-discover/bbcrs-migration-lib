package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReferencePointTranslation Reference Point Translation
type ReferencePointTranslation struct {
	basic.Base
	basic.DataOwner
	ReferencePointTranslationAPI
	ReferencePointID *uuid.UUID `json:"reference_point_id,omitempty" gorm:"type:varchar(36);uniqueIndex:reference_point_translation_unique;not null" swaggertype:"string" format:"uuid"` // Reference Point id
}

// ReferencePointTranslationAPI Reference Point Translation API
type ReferencePointTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:reference_point_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ReferencePointName *string `json:"reference_point_name,omitempty" gorm:"type:varchar(256)"`                                                             // Reference Point Name
	Description        *string `json:"description,omitempty" gorm:"type:text"`                                                                              // Description
}
