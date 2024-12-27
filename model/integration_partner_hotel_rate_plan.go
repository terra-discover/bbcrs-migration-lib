package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelRatePlan Integration Partner Hotel Rate Plan
type IntegrationPartnerHotelRatePlan struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelRatePlanAPI
	IntegrationPartnerHotelRatePlanTranslation *IntegrationPartnerHotelRatePlanTranslation `json:"integration_partner_hotel_rate_plan_translation,omitempty"` // Integration Partner Hotel Rate Plan Translation
	IntegrationPartnerHotel                    *IntegrationPartnerHotel                    `json:"integration_partner_hotel,omitempty"`                       // Integration Partner Hotel
}

// IntegrationPartnerHotelRatePlanAPI Integration Partner Hotel Rate Plan API
type IntegrationPartnerHotelRatePlanAPI struct {
	IntegrationPartnerHotelID *uuid.UUID `json:"integration_partner_hotel_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_rate_plan__rate_plan_code,unique,where:deleted_at is null;not null"` // Integration Partner Hotel ID
	RatePlanID                *uuid.UUID `json:"rate_plan_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_hotel_rate_plan__rate_plan_code,unique,where:deleted_at is null;not null"`                 // Rate Plan ID
	RatePlanCode              *string    `json:"rate_plan_code,omitempty" gorm:"type:varchar(256);index:idx_integration_partner_hotel_rate_plan__rate_plan_code,unique,where:deleted_at is null;not null"`                                                 // Rate Plan Code
	RatePlanName              *string    `json:"rate_plan_name,omitempty" gorm:"type:varchar(512)"`                                                                                                                                                        // Rate Plan Name
	RateComments              *string    `json:"rate_comments,omitempty" gorm:"type:text"`                                                                                                                                                                 // Rate Comments
	RateCommentDetails        *string    `json:"rate_comment_details,omitempty" gorm:"type:text"`                                                                                                                                                          // Rate Comment Details
	Description               *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                   // Description
	Comment                   *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                       // Comment
}
