package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerFeeTaxType Integration Partner Fee Tax Type
type IntegrationPartnerFeeTaxType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerFeeTaxTypeAPI
	FeeTaxType              *FeeTaxType              `json:"fee_tax_type,omitempty" gorm:"foreignKey:FeeTaxTypeID;references:ID"`
	AgentIntegrationPartner *AgentIntegrationPartner `json:"agent_integration_partner,omitempty" gorm:"foreignKey:IntegrationPartnerID;references:IntegrationPartnerID"`
	IntegrationPartner      *IntegrationPartner      `json:"integration_partner,omitempty" gorm:"foreignKey:IntegrationPartnerID;references:ID"`
}

// IntegrationPartnerFeeTaxTypeAPI Integration Partner Fee Tax Type API
type IntegrationPartnerFeeTaxTypeAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string"`
	FeeTaxTypeID         *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	FeeTaxTypeCode       *string    `json:"fee_tax_type_code,omitempty" gorm:"type:varchar(36)" example:"op"`
	FeeTaxTypeName       *string    `json:"fee_tax_type_name,omitempty" gorm:"type:varchar(256)" example:"aviation levy"`
	Description          *string    `json:"description,omitempty" gorm:"type:text"`
}

// IntegrationPartnerFeeTaxTypeDataGet struct
type IntegrationPartnerFeeTaxTypeDataGet struct {
	IntegrationPartnerFeeTaxTypeDataGetAPI
	ID                             *uuid.UUID       `json:"id,omitempty"`
	IntegrationPartnerID           *uuid.UUID       `json:"-"`
	IntegrationPartnerFeeTaxTypeID *uuid.UUID       `json:"-"`
	FeeTaxTypeID                   *uuid.UUID       `json:"-"`
	CreatedAt                      *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt                      *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort                           *int64           `json:"sort,omitempty"`                                               // sort
	Status                         *int64           `json:"status,omitempty"`
}

// IntegrationPartnerFeeTaxTypeDataGetAPI struct
type IntegrationPartnerFeeTaxTypeDataGetAPI struct {
	FeeTaxType                   *FeeTaxType                   `json:"FeeTaxType" gorm:"foreignKey:FeeTaxTypeID;references:ID"`
	IntegrationPartner           *IntegrationPartner           `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"`
	IntegrationPartnerFeeTaxType *IntegrationPartnerFeeTaxType `json:"integration_partner_fee_tax_type" gorm:"foreignKey:IntegrationPartnerFeeTaxTypeID;references:ID"`
}
