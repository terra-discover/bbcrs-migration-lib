package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerSegmentCategoryTranslation Integration Partner Segment Category Translation
type IntegrationPartnerSegmentCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerSegmentCategoryTranslationAPI
	IntegrationPartnerSegmentCategoryID *uuid.UUID `json:"integration_partner_segment_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_segment_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Segment Category id
}

// IntegrationPartnerSegmentCategoryTranslationAPI Integration Partner Segment Category Translation API
type IntegrationPartnerSegmentCategoryTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_segment_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SegmentCategoryName *string `json:"segment_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Segment Category Name
	Description         *string `json:"description,omitempty" gorm:"type:text"`                                                                                                   // Description
	Comment             *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                       // Comment
}
