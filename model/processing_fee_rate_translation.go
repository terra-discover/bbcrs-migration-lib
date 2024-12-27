package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateTranslation Processing Fee Rate Translation
type ProcessingFeeRateTranslation struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateTranslationAPI
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);uniqueIndex:processing_fee_rate_translation_unique;not null" swaggertype:"string" format:"uuid"` // Processing Fee Rate id
}

// ProcessingFeeRateTranslationAPI Processing Fee Rate Translation API
type ProcessingFeeRateTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:processing_fee_rate_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                                  // Description
}
