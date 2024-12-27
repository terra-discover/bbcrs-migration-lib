package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// EmailAddressTypeTranslation EmailAddressType Translation
type EmailAddressTypeTranslation struct {
	basic.Base
	basic.DataOwner
	EmailAddressTypeTranslationAPI
	EmailAddressTypeID *uuid.UUID `json:"email_address_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:email_address_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Email Address Type id
}

// EmailAddressTypeTranslationAPI EmailAddressType Translation API
type EmailAddressTypeTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:email_address_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	EmailAddressTypeName *string `json:"email_address_type_name,omitempty" gorm:"type:varchar(256);"`                                                            // Email Address Type Name
}
