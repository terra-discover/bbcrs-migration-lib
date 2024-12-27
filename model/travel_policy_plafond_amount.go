package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TravelPolicyPlafondAmount Travel Policy Job Title
type TravelPolicyPlafondAmount struct {
	basic.Base
	basic.DataOwner
	TravelPolicyPlafondAmountAPI
	TravelPolicy *TravelPolicy `json:"travel_policy,omitempty" gorm:"foreignKey:TravelPolicyID;references:ID"` // travel policy
	Currency     *Currency     `json:"currency,omitempty" gorm:"foreignKey:CurrencyID;reference:ID"`           // currency
}

// TravelPolicyPlafondAmountAPI Travel Policy Job Title API
type TravelPolicyPlafondAmountAPI struct {
	TravelPolicyID *uuid.UUID `json:"travel_policy_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"` // travel policy id
	CurrencyID     *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid"`      // currency id
	IsTaxInclusive *bool      `json:"is_tax_inclusive,omitempty" example:"true"`
	Amount         *float64   `json:"amount,omitempty"`
}
