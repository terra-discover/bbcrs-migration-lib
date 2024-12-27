package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerLocationCategoryTranslation Integration Partner Location Category Translation
type IntegrationPartnerLocationCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerLocationCategoryTranslationAPI
	IntegrationPartnerLocationCategoryID *uuid.UUID `json:"integration_partner_location_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_location_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Location Category id
}

// IntegrationPartnerLocationCategoryTranslationAPI Integration Partner Location Category Translation API
type IntegrationPartnerLocationCategoryTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_location_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	LocationCategoryName *string `json:"location_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Location Category Name
	Description          *string `json:"description,omitempty" gorm:"type:text"`                                                                                                    // Description
	Comment              *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                        // Comment
}
