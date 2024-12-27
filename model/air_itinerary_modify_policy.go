package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Air Itinerary Modify Policy
type AirItineraryModifyPolicy struct {
	basic.Base
	basic.DataOwner
	AirItineraryModifyPolicyAPI
}

// AirItineraryModifyPolicyAPI
type AirItineraryModifyPolicyAPI struct {
	AirItineraryID        *uuid.UUID `json:"air_itinerary_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ModifyPolicyID        *uuid.UUID `json:"modify_policy_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ModifyPolicyCode      *string    `json:"modify_policy_code,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	ModifyPolicyName      *string    `json:"modify_policy_name,omitempty" gorm:"type:varchar(256)" swaggertype:"string"`
	OffsetTimeUnitID      *uuid.UUID `json:"offset_time_unit_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	OffsetUnitMultiplier  *int       `json:"offset_unit_multiplier,omitempty" gorm:"smallint"`
	OffsetDropTimeID      *uuid.UUID `json:"offset_drop_time_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	IsTaxInclusive        *bool      `json:"is_tax_inclusive,omitempty"`
	IsServiceFeeInclusive *bool      `json:"is_service_fee_inclusive,omitempty"`
	IsFeeInclusive        *bool      `json:"is_fee_inclusive,omitempty"`
	NotModifiable         *bool      `json:"not_modifiable,omitempty"`
	NotRefundable         *bool      `json:"not_refundable,omitempty"`
	CurrencyID            *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	Amount                *float64   `json:"amount,omitempty"`
	Description           *string    `json:"description,omitempty" gorm:"text"`
}
