package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerPromotion Integration Partner Promotion
type IntegrationPartnerPromotion struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerPromotionAPI
	IntegrationPartnerPromotionTranslation *IntegrationPartnerPromotionTranslation `json:"integration_partner_promotion_translation,omitempty"` // Integration Partner Promotion Translation
}

// IntegrationPartnerPromotionAPI Integration Partner Promotion API
type IntegrationPartnerPromotionAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_promotion__promotion_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	PromotionID          *uuid.UUID `json:"promotion_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_promotion__promotion_code,unique,where:deleted_at is null;not null"`           // Promotion ID
	PromotionCode        *string    `json:"promotion_code,omitempty" gorm:"type:varchar(36);index:idx_integration_partner_promotion__promotion_code,unique,where:deleted_at is null;not null"`                                            // Promotion Code
	PromotionName        *string    `json:"promotion_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                            // Promotion Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                       // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                           // Comment
}
