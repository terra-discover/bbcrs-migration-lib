package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerSegmentCategory Integration Partner Segment Category
type IntegrationPartnerSegmentCategory struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerSegmentCategoryAPI
	IntegrationPartnerSegmentCategoryTranslation *IntegrationPartnerSegmentCategoryTranslation `json:"integration_partner_segment_category_translation,omitempty"` // Integration Partner Segment Category Translation
	SegmentCategory                              *SegmentCategory                              `json:"segment_category,omitempty"`                                 // Segment Category
}

// IntegrationPartnerSegmentCategoryAPI Integration Partner Segment Category API
type IntegrationPartnerSegmentCategoryAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_segment_category__segment_category_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	SegmentCategoryID    *uuid.UUID `json:"segment_category_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_segment_category__segment_category_code,unique,where:deleted_at is null;not null"`    // Segment Category ID
	SegmentCategoryCode  *string    `json:"segment_category_code,omitempty" gorm:"type:varchar(8);index:idx_integration_partner_segment_category__segment_category_code,unique,where:deleted_at is null;not null"`                                      // Segment Category Code
	SegmentCategoryName  *string    `json:"segment_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                   // Segment Category Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                     // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                         // Comment
}
