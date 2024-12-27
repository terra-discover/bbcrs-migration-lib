package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoundNumberFunctionTranslation Round Number Function Translation
type RoundNumberFunctionTranslation struct {
	basic.Base
	basic.DataOwner
	RoundNumberFunctionTranslationAPI
	RoundNumberFunctionID *uuid.UUID `json:"round_number_function_id,omitempty" gorm:"type:varchar(36);uniqueIndex:round_number_function_translation_unique;not null" swaggertype:"string" format:"uuid"` // Round Number Function id
}

// RoundNumberFunctionTranslationAPI Round Number Function Translation API
type RoundNumberFunctionTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:round_number_function_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RoundFunctionName *string `json:"round_function_name,omitempty" gorm:"type:varchar(256)"`                                                                    // Round Function Name
}
