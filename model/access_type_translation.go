package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AccessTypeTranslation AccessType Translation
type AccessTypeTranslation struct {
	basic.Base
	basic.DataOwner
	AccessTypeTranslationAPI
	AccessTypeID *uuid.UUID `json:"access_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:idx__access_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Agent User Type ID
}

// AccessTypeTranslationAPI AccessType Translation API
type AccessTypeTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:idx__access_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...                                                                   // Description
	AccessTypeName *string `json:"access_type_name,omitempty" gorm:"type:varchar(256)" example:"Admin"`
}
