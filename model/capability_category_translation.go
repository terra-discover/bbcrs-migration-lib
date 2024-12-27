package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CapabilityCategoryTranslation Capability Category Translation
type CapabilityCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	CapabilityCategoryTranslationAPI
	CapabilityCategoryID *uuid.UUID `json:"capability_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:capability_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Capability Category id
}

// CapabilityCategoryTranslationAPI Capability Category Translation API
type CapabilityCategoryTranslationAPI struct {
	LanguageCode           *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:capability_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CapabilityCategoryName *string `json:"capability_category_name,omitempty" gorm:"type:varchar(256)"`                                                             // Capability Category Name
	Description            *string `json:"description,omitempty" gorm:"type:text"`                                                                                  // Description
}
