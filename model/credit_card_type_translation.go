package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CreditCardTypeTranslation Credit Card Type Translation
type CreditCardTypeTranslation struct {
	basic.Base
	basic.DataOwner
	CreditCardTypeTranslationAPI
	CreditCardTypeID *uuid.UUID `json:"credit_card_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:credit_card_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Credit Card Type id
}

// CreditCardTypeTranslationAPI Credit Card Type Translation API
type CreditCardTypeTranslationAPI struct {
	LanguageCode             *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:credit_card_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CreditCardTypeName       *string `json:"credit_card_type_name,omitempty" gorm:"type:varchar(256)" example:"bank"`                                              // Credit Card Type Name
	CardIdentificationRemark *string `json:"card_identification_remark,omitempty" gorm:"type:text" example:"12345"`                                                // Card Identification Remark
}
