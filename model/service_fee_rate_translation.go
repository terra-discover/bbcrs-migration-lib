package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateTranslation Task Type Translation
type ServiceFeeRateTranslation struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateTranslationAPI
	ServiceFeeRate   *ServiceFeeRate `json:"service_fee_rate,omitempty" swaggerignore:"true"`
	Language         *Language       `json:"language,omitempty"`
	ServiceFeeRateID *uuid.UUID      `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);uniqueIndex:service_fee_rate_translation_unique;not null" swaggertype:"string" format:"uuid"` // ServiceFee Rate id
}

// ServiceFeeRateTranslationAPI Task Type Translation API
type ServiceFeeRateTranslationAPI struct {
	LanguageID   *uuid.UUID `json:"language_id,omitempty" gorm:"type:varchar(36);uniqueIndex:service_fee_rate_id_translation_unique;not null" swaggertype:"string" format:"uuid"` // Language ID
	LanguageCode *string    `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:service_fee_rate_code_translation_unique;not null" example:"en"`                    // Language code example: en, id, cn etc...
	Description  *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                       // Description
}
