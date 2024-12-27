package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PhoneUseTypeTranslation Phone Use Type Translation
type PhoneUseTypeTranslation struct {
	basic.Base
	basic.DataOwner
	PhoneUseTypeTranslationAPI
	PhoneUseTypeID *uuid.UUID `json:"phone_use_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:phone_use_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Phone Use Type id
}

// PhoneUseTypeTranslationAPI Phone Use Type Translation API
type PhoneUseTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:phone_use_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PhoneUseTypeName *string `json:"phone_use_type_name,omitempty" gorm:"type:varchar(256)" example:"Contact"`                                           // Phone Use Type Name
}
