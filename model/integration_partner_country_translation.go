package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCountryTranslation Integration Partner Country Translation
type IntegrationPartnerCountryTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCountryTranslationAPI
	IntegrationPartnerCountryID *uuid.UUID `json:"integration_partner_country_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_country_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Country id
}

// IntegrationPartnerCountryTranslationAPI Integration Partner Country Translation API
type IntegrationPartnerCountryTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_country_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CountryName  *string `json:"country_name,omitempty" gorm:"type:varchar(64)"`                                                                                  // Country Name
}
