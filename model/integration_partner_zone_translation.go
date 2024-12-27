package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerZoneTranslation Integration Partner Zone Translation
type IntegrationPartnerZoneTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerZoneTranslationAPI
	IntegrationPartnerZoneID *uuid.UUID `json:"integration_partner_zone_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_zone_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Zone id
}

// IntegrationPartnerZoneTranslationAPI Integration Partner Zone Translation API
type IntegrationPartnerZoneTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_zone_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	ZoneName     *string `json:"zone_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Zone Name
	Description  *string `json:"description,omitempty" gorm:"type:text"`                                                                                       // Description
}
