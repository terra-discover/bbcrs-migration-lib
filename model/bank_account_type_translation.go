package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BankAccountTypeTranslation Bank Account Type Translation
type BankAccountTypeTranslation struct {
	basic.Base
	basic.DataOwner
	BankAccountTypeTranslationAPI
	BankAccountTypeID *uuid.UUID `json:"bank_account_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:bank_account_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Bank Account Type id
}

// BankAccountTypeTranslationAPI Bank Account Type Translation API
type BankAccountTypeTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:bank_account_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	BankAccountTypeName *string `json:"bank_account_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Bank Account Type Name
}
