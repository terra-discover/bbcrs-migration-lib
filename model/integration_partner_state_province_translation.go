package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerStateProvinceTranslation Integration Partner State Province Translation
type IntegrationPartnerStateProvinceTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerStateProvinceTranslationAPI
	IntegrationPartnerStateProvinceID *uuid.UUID `json:"integration_partner_state_province_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_state_province_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner State Province id
}

// IntegrationPartnerStateProvinceTranslationAPI Integration Partner State Province Translation API
type IntegrationPartnerStateProvinceTranslationAPI struct {
	LanguageCode      *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_state_province_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	StateProvinceName *string `json:"state_province_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // State Province Name
	Description       *string `json:"description,omitempty" gorm:"type:text"`                                                                                                 // Description
	Comment           *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                     // Comment
}
