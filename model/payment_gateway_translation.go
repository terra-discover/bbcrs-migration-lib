package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PaymentGatewayTranslation Payment Gateway Translation
type PaymentGatewayTranslation struct {
	basic.Base
	basic.DataOwner
	PaymentGatewayTranslationAPI
	PaymentGatewayID *uuid.UUID `json:"payment_gateway_id,omitempty" gorm:"type:varchar(36);uniqueIndex:payment_gateway_translation_unique;not null" swaggertype:"string" format:"uuid"` // Payment Gateway id
}

// PaymentGatewayTranslationAPI Payment Gateway Translation API
type PaymentGatewayTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:payment_gateway_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PaymentGatewayName *string `json:"payment_gateway_name,omitempty" example:"Credit Card" gorm:"type:varchar(256)"`                                       // Payment Gateway Name                                                                             // Description
}
