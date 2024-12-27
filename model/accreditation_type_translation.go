package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AccreditationTypeTranslation Accreditation Type Translation
type AccreditationTypeTranslation struct {
	basic.Base
	basic.DataOwner
	AccreditationTypeTranslationAPI
	AccreditationTypeID *uuid.UUID `json:"accreditation_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:accreditation_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Accreditation Type id
}

// AccreditationTypeTranslationAPI Accreditation Type Translation API
type AccreditationTypeTranslationAPI struct {
	LanguageCode          *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:accreditation_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AccreditationTypeName *string `json:"accreditation_type_name,omitempty" gorm:"type:varchar(256)" example:"International Air Transport Association (IATA)"`    // Accreditation Type Name
}
