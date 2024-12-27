package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoundNumberDigitTranslation Round Number Digit Translation
type RoundNumberDigitTranslation struct {
	basic.Base
	basic.DataOwner
	RoundNumberDigitTranslationAPI
	RoundNumberDigitID *uuid.UUID `json:"round_number_digit_id,omitempty" gorm:"type:varchar(36);uniqueIndex:round_number_digit_translation_unique;not null" swaggertype:"string" format:"uuid"` // Round Number Digit id
}

// RoundNumberDigitTranslationAPI Round Number Digit Translation API
type RoundNumberDigitTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:round_number_digit_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoundDigitName *string `json:"round_digit_name,omitempty" gorm:"type:varchar(256)"`                                                                    // Round Digit Name
}
