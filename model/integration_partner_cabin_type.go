package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCabinType Integration Partner Cabin Type
type IntegrationPartnerCabinType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCabinTypeAPI
	IntegrationPartner *IntegrationPartner `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
	CabinType          *CabinType          `json:"cabin_type" gorm:"foreignKey:CabinTypeID;references:ID"`                   // Cabin Type
}

// IntegrationPartnerCabinTypeAPI Integration Partner Cabin Type API
type IntegrationPartnerCabinTypeAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid"`                         // Integration Partner ID
	CabinTypeID          *uuid.UUID `json:"cabin_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null"` // Cabin Type ID
	CabinTypeCode        *string    `json:"cabin_type_code,omitempty" gorm:"type:varchar(36);not null"`                                  // Cabin Type Code
	CabinTypeName        *string    `json:"cabin_type_name,omitempty" gorm:"type:varchar(256);not null"`                                 // Cabin Type Name
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                          // Comment
}
