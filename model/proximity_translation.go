package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProximityTranslation Proximity Translation
type ProximityTranslation struct {
	basic.Base
	basic.DataOwner
	ProximityTranslationAPI
	ProximityID *uuid.UUID `json:"proximity_id,omitempty" gorm:"type:varchar(36);uniqueIndex:proximity_translation_unique;not null" swaggertype:"string" format:"uuid"` // Proximity id
}

// ProximityTranslationAPI Proximity Translation API
type ProximityTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:proximity_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ProximityName *string `json:"proximity_name,omitempty" gorm:"type:varchar(256)" example:"Onsite"`                                            // Proximity Name
}
