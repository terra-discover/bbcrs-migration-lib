package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CommissionPolicy Commission Policy
type CommissionPolicy struct {
	basic.Base
	basic.DataOwner
	CommissionPolicyAPI
	CommissionPolicyTranslation *CommissionPolicyTranslation `json:"commission_policy_translation,omitempty"` // Commission Policy Translation
	ChargeType                  *ChargeType                  `json:"charge_type,omitempty"`                   // Charge Type
	Currency                    *Currency                    `json:"currency,omitempty"`                      // Currency
}

// CommissionPolicyAPI Commission Policy API
type CommissionPolicyAPI struct {
	CommissionPolicyCode       *string    `json:"commission_policy_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`  // Commission Policy Code
	CommissionPolicyName       *string    `json:"commission_policy_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Commission Policy Name
	ChargeTypeID               *uuid.UUID `json:"charge_type_id,omitempty" swaggertype:"string" format:"uuid"`                                               // Charge Type ID
	Percent                    *float64   `json:"percent,omitempty"`                                                                                         // Percent
	CurrencyID                 *uuid.UUID `json:"currency_id,omitempty" swaggertype:"string" format:"uuid"`                                                  // Currency ID
	FlatCommission             *float64   `json:"flat_commission,omitempty"`                                                                                 // Flat Commission
	IsTaxInclusive             *bool      `json:"is_tax_inclusive,omitempty"`                                                                                // Is Tax Inclusive
	IsFeeInclusive             *bool      `json:"is_fee_inclusive,omitempty"`                                                                                // Is Fee Inclusive
	IsServiceFeeInclusive      *bool      `json:"is_service_fee_inclusive,omitempty"`                                                                        // Is Service Fee Inclusive
	IsSellableProductInclusive *bool      `json:"is_sellable_product_inclusive,omitempty"`                                                                   // Is Sellable Product Inclusive
	Description                *string    `json:"description,omitempty" gorm:"type:text"`                                                                    // Description
}
