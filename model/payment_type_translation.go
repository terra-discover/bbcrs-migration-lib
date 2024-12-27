package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PaymentTypeTranslation Payment Type Translation
type PaymentTypeTranslation struct {
	basic.Base
	basic.DataOwner
	PaymentTypeTranslationAPI
	PaymentTypeID *uuid.UUID `json:"payment_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:payment_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Payment Type id
}

// PaymentTypeTranslationAPI Payment Type Translation API
type PaymentTypeTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:payment_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PaymentTypeName *string `json:"payment_type_name,omitempty" gorm:"type:varchar(256)" example:"Bank Transfer"`                                     // Payment Type Name
}
