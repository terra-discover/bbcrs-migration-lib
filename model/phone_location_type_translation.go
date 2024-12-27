package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PhoneLocationTypeTranslation Phone Location Type Translation
type PhoneLocationTypeTranslation struct {
	basic.Base
	basic.DataOwner
	PhoneLocationTypeTranslationAPI
	PhoneLocationTypeID *uuid.UUID `json:"phone_location_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:phone_location_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Phone Location Type id
}

// PhoneLocationTypeTranslationAPI Phone Location Type Translation API
type PhoneLocationTypeTranslationAPI struct {
	LanguageCode          *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:phone_location_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PhoneLocationTypeName *string `json:"phone_location_type_name,omitempty" gorm:"type:varchar(256)" example:"Home"`                                              // Phone Location Type Name
}
