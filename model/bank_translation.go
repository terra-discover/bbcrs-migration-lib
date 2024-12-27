package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// BankTranslation Bank Translation
type BankTranslation struct {
	basic.Base
	basic.DataOwner
	BankTranslationAPI
	BankID *uuid.UUID `json:"bank_id,omitempty" gorm:"type:varchar(36);uniqueIndex:bank_translation_unique;not null" swaggertype:"string" format:"uuid"` // Bank id
}

// BankTranslationAPI Bank Translation API
type BankTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:bank_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	BankName     *string `json:"bank_name,omitempty" gorm:"type:varchar(256)" example:"bank"`                                              // Bank Name
}
