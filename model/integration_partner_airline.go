package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerAirline Integration Partner Airline
type IntegrationPartnerAirline struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerAirlineAPI
	IntegrationPartner *IntegrationPartner `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
	Airline            *Airline            `json:"airline" gorm:"foreignKey:AirlineID;references:ID"`                        // Airline
	Currency           *Currency           `json:"currency" gorm:"foreignKey:CurrencyID;references:ID"`                      // Currency
}

// IntegrationPartnerAirlineAPI Integration Partner Airline API
type IntegrationPartnerAirlineAPI struct {
	IntegrationPartnerID       *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Integration Partner ID
	AirlineID                  *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`             // Airline ID
	WebServiceUserName         *string    `json:"web_service_user_name,omitempty" gorm:"type:varchar(256)"`                                             // Web Service User Name
	WebServicePassword         *string    `json:"web_service_password,omitempty" gorm:"type:varchar(256)"`                                              // Web Service Password
	PartnerAirlineCode         *string    `json:"partner_airline_code,omitempty" gorm:"type:varchar(36)"`                                               // Partner Airline Code
	PartnerAirlineName         *string    `json:"partner_airline_name,omitempty" gorm:"type:varchar(64)"`                                               // Partner Airline Name
	CurrencyID                 *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                     // Currency ID
	ExtranetUrl                *string    `json:"extranet_url,omitempty" gorm:"type:varchar(256)"`                                                      // Extranet Url
	ExtranetUserName           *string    `json:"extranet_user_name,omitempty" gorm:"type:varchar(256)"`                                                // Extranet User Name
	ExtranetPassword           *string    `json:"extranet_password,omitempty" gorm:"type:varchar(256)"`                                                 // Extranet Password
	IsEnabled                  *bool      `json:"is_enabled,omitempty"`                                                                                 // Is Enabled
	IsPrimary                  *bool      `json:"is_primary,omitempty"`                                                                                 // Is Primary
	NeedNationalIdentityNumber *bool      `json:"need_national_identity_number,omitempty"`                                                              // Need National Identity Number
	Comment                    *string    `json:"comment,omitempty" gorm:"type:text"`                                                                   // Comment
}
