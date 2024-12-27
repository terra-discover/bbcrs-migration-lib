package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateServiceFeeCategory Model
type CorporateServiceFeeCategory struct {
	basic.Base
	basic.DataOwner
	CorporateID *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Corporate ID
	CorporateServiceFeeCategoryAPI
	Corporate          *Corporate          `json:"corporate,omitempty"`                                                                                                    // Corporate
	ServiceFeeCategory *ServiceFeeCategory `json:"service_fee_category,omitempty"`                                                                                         // Service Fee Category
	ServiceFeeRate     *ServiceFeeRate     `json:"service_fee_rate,omitempty" gorm:"foreignKey:ServiceFeeCategoryID;references:ServiceFeeCategoryID" swaggerignore:"true"` // Service Fee Rate
	ServiceLevel       *ServiceLevel       `json:"service_level,omitempty"`                                                                                                // Service Level
}

// CorporateServiceFeeCategoryAPI Model
type CorporateServiceFeeCategoryAPI struct {
	ServiceFeeCategoryID *uuid.UUID `json:"service_fee_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // ServiceFee Category ID
	ServiceLevelID       *uuid.UUID `json:"service_level_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                 // Service Level ID
	IsSelfBookingTool    *bool      `json:"is_self_booking_tool,omitempty" gorm:"default:false"`                                                   // Is Self Booking Tool
	IsPersonalTravel     *bool      `json:"is_personal_travel,omitempty" gorm:"default:false"`                                                     // Is Personal Travel
	IsIncluded           *bool      `json:"is_included,omitempty"`                                                                                 // Is Included
	IsHidden             *bool      `json:"is_hidden,omitempty"`                                                                                   // Is Hidden
}
