package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PhoneTechnologyTypeTranslation Phone Technology Type Translation
type PhoneTechnologyTypeTranslation struct {
	basic.Base
	basic.DataOwner
	PhoneTechnologyTypeTranslationAPI
	PhoneTechnologyTypeID *uuid.UUID `json:"phone_technology_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:phone_technology_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Phone Technology Type id
}

// PhoneTechnologyTypeTranslationAPI Phone Technology Type Translation API
type PhoneTechnologyTypeTranslationAPI struct {
	LanguageCode            *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:phone_technology_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PhoneTechnologyTypeName *string `json:"phone_technology_type_name,omitempty" gorm:"type:varchar(256)" example:"Voice"`                                             // Phone Technology Type Name
}
