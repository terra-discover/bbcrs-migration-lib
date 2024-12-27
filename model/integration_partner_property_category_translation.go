package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerPropertyCategoryTranslation Integration Partner Property Category Translation
type IntegrationPartnerPropertyCategoryTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerPropertyCategoryTranslationAPI
	IntegrationPartnerPropertyCategoryID *uuid.UUID `json:"integration_partner_property_category_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_property_category_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Property Category id
}

// IntegrationPartnerPropertyCategoryTranslationAPI Integration Partner Property Category Translation API
type IntegrationPartnerPropertyCategoryTranslationAPI struct {
	LanguageCode         *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_property_category_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	PropertyCategoryName *string `json:"property_category_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Property Category Name
	Description          *string `json:"description,omitempty" gorm:"type:text"`                                                                                                    // Description
	Comment              *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                        // Comment
}
