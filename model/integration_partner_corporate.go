package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCorporate Integration Partner Corporate
type IntegrationPartnerCorporate struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCorporateAPI
	Corporate          *Corporate          `json:"corporate" gorm:"foreignKey:CorporateID;references:ID"`                    // Corporate
	IntegrationPartner *IntegrationPartner `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
}

// IntegrationPartnerCorporateAPI Integration Partner Corporate API
type IntegrationPartnerCorporateAPI struct {
	IntegrationPartnerID        *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Integration Partner ID
	CorporateID                 *uuid.UUID `json:"corporate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`           // Corporate ID
	CorporateCode               *string    `json:"corporate_code,omitempty" gorm:"type:varchar(36)"`                                                     // Partner's internal corporate code
	CorporateInformationEnabled *bool      `json:"corporate_information_enabled,omitempty"`                                                              // Corporate Information Enabled
	CorporatePerformanceEnabled *bool      `json:"corporate_performance_enabled,omitempty"`                                                              // Corporate Performance Enabled
	CreditLimitEnabled          *bool      `json:"credit_limit_enabled,omitempty"`                                                                       // Credit Limit Enabled
	IsEnabled                   *bool      `json:"is_enabled,omitempty"`                                                                                 // Is Enabled
}
