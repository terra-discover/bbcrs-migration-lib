package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCityTranslation Integration Partner City Translation
type IntegrationPartnerCityTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCityTranslationAPI
	IntegrationPartnerCityID *uuid.UUID `json:"integration_partner_city_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_city_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner City id
}

// IntegrationPartnerCityTranslationAPI Integration Partner City Translation API
type IntegrationPartnerCityTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_city_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CityName     *string `json:"city_name,omitempty" gorm:"type:varchar(64)"`                                                                                  // City Name
}
