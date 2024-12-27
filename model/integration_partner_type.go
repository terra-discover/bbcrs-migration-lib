package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// IntegrationPartnerType struct
type IntegrationPartnerType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerTypeAPI
}

// IntegrationPartnerTypeAPI struct
type IntegrationPartnerTypeAPI struct {
	IntegrationPartnerTypeCode *int64  `json:"integration_partner_type_code,omitempty" gorm:"not null"`                                                          // IntegrationPartnerTypeCode
	IntegrationPartnerTypeName *string `json:"integration_partner_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // IntegrationPartnerTypeName
}
