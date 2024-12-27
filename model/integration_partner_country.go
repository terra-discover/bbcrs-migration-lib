package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCountry Integration Partner Country
type IntegrationPartnerCountry struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCountryAPI
	IntegrationPartnerCountryTranslation *IntegrationPartnerCountryTranslation `json:"integration_partner_country_translation,omitempty"` // Integration Partner Country Translation
	Country                              *Country                              `json:"country,omitempty"`                                 // Country
}

// IntegrationPartnerCountryAPI Integration Partner Country API
type IntegrationPartnerCountryAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"index:idx_idx_integration_partner_country__country_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	CountryID            *uuid.UUID `json:"country_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                                                     // Country ID
	CountryCode          *string    `json:"country_code,omitempty" gorm:"type:varchar(36);index:idx_idx_integration_partner_country__country_code,unique,where:deleted_at is null;not null"`                             // Country Code
	CountryName          *string    `json:"country_name,omitempty" gorm:"type:varchar(64)"`                                                                                                                              // Country Name
	Nationality          *string    `json:"nationality,omitempty" gorm:"type:varchar(64)"`                                                                                                                               // Nationality
}
