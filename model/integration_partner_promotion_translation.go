package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerPromotionTranslation Integration Partner Promotion Translation
type IntegrationPartnerPromotionTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerPromotionTranslationAPI
	IntegrationPartnerPromotionID *uuid.UUID `json:"integration_partner_promotion_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_promotion_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Promotion id
}

// IntegrationPartnerPromotionTranslationAPI Integration Partner Promotion Translation API
type IntegrationPartnerPromotionTranslationAPI struct {
	LanguageCode  *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_promotion_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PromotionName *string `json:"promotion_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Promotion Name
	Description   *string `json:"description,omitempty" gorm:"type:text"`                                                                                            // Description
	Comment       *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                // Comment
}
