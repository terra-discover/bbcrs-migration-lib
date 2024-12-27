package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AddressTypeTranslation Address Type Translation
type AddressTypeTranslation struct {
	basic.Base
	basic.DataOwner
	AddressTypeTranslationAPI
	AddressTypeID *uuid.UUID `json:"address_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:address_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Address Type id
}

// AddressTypeTranslationAPI Address Type Translation API
type AddressTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:address_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AddressTypeName *string `json:"address_type_name,omitempty" gorm:"type:varchar(256)" example:"Delivery"`                                          // Address Type Name
}
