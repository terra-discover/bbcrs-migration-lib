package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelRatePlanTranslation Integration Partner Hotel Rate Plan Translation
type IntegrationPartnerHotelRatePlanTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelRatePlanTranslationAPI
	IntegrationPartnerHotelRatePlanID *uuid.UUID `json:"integration_partner_hotel_rate_plan_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_hotel_rate_plan_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Hotel Rate Plan id
}

// IntegrationPartnerHotelRatePlanTranslationAPI Integration Partner Hotel Rate Plan Translation API
type IntegrationPartnerHotelRatePlanTranslationAPI struct {
	LanguageCode       *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_hotel_rate_plan_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RatePlanName       *string `json:"rate_plan_name,omitempty" gorm:"type:varchar(512)"`                                                                                       // Rate Plan Name
	RateComments       *string `json:"rate_comments,omitempty" gorm:"type:text"`                                                                                                // Rate Comments
	RateCommentDetails *string `json:"rate_comment_details,omitempty" gorm:"type:text"`                                                                                         // Rate Comment Details
	Description        *string `json:"description,omitempty" gorm:"type:text"`                                                                                                  // Description
	Comment            *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                      // Comment
}
