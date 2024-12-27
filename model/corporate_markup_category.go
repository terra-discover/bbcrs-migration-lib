package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateMarkupCategory Model
type CorporateMarkupCategory struct {
	basic.Base
	basic.DataOwner
	CorporateMarkupCategoryAPI
	Corporate      *Corporate      `json:"corporate,omitempty"`       // Corporate
	MarkupCategory *MarkupCategory `json:"markup_category,omitempty"` // Markup Category
	ServiceLevel   *ServiceLevel   `json:"service_level,omitempty"`   // Service Level
}

// CorporateMarkupCategoryAPI Model
type CorporateMarkupCategoryAPI struct {
	CorporateID       *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`       // Corporate ID
	MarkupCategoryID  *uuid.UUID `json:"markup_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Markup Category ID
	ServiceLevelID    *uuid.UUID `json:"service_level_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Service Level ID
	IsSelfBookingTool *bool      `json:"is_self_booking_tool,omitempty" gorm:"default:false"`                                              // Is Self Booking Tool
	IsPersonalTravel  *bool      `json:"is_personal_travel,omitempty" gorm:"default:false"`                                                // Is Personal Travel
}
