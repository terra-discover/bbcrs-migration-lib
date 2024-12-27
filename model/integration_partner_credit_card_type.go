package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCreditCardType Integration Partner Credit Card Type
type IntegrationPartnerCreditCardType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCreditCardTypeAPI
}

// IntegrationPartnerCreditCardTypeAPI Integration PartnerCredit Card Type API
type IntegrationPartnerCreditCardTypeAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`
	CreditCardTypeID     *uuid.UUID `json:"credit_card_type_id,omitempty" gorm:"type:varchar(36);not null;"`
	CreditCardTypeCode   *string    `json:"credit_card_type_code,omitempty" gorm:"type:varchar(64);index:,unique,where:deleted_at is null;not null" example:"bnk"`   // Credit Card Type Code
	CreditCardTypeName   *string    `json:"credit_card_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"bank"` // Credit Card Type Name
}
