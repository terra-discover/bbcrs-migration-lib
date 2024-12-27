package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirItineraryGuaranteePaymentPolicy struct {
	basic.Base
	basic.DataOwner
	AirItineraryGuaranteePaymentPolicyAPI
}

type AirItineraryGuaranteePaymentPolicyAPI struct {
	AirItineraryID             *uuid.UUID       `json:"air_itinerary_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	GuaranteePaymentPolicyID   *uuid.UUID       `json:"guarantee_payment_policy_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	GuaranteePaymentPolicyCode *string          `json:"guarantee_payment_policy_code,omitempty" gorm:"type:varchar(36)"`
	GuaranteePaymentPolicyName *string          `json:"guarantee_payment_policy_name,omitempty" gorm:"type:varchar(256)"`
	GuaranteeTypeID            *uuid.UUID       `json:"guarantee_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	PaymentTypeID              *uuid.UUID       `json:"payment_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	OffsetTimeUnitID           *string          `json:"offset_time_unit_id,omitempty" gorm:"type:varchar(36)"`
	BasisTypeID                *string          `json:"basis_type_id,omitempty" gorm:"type:varchar(36)"`
	IsTaxInclusive             *bool            `json:"is_tax_inclusive,omitempty"`
	IsServiceFeeInclusive      *bool            `json:"is_service_fee_inclusive,omitempty"`
	IsFeeInclusive             *bool            `json:"is_fee_inclusive,omitempty"`
	AbsoluteDeadline           *strfmt.DateTime `json:"absolute_deadline,omitempty" gorm:"type:timestamptz" format:"date-time"`
	OffsetUnitMultiplier       *int             `json:"offset_unit_multiplier,omitempty" gorm:"type:smallint"`
	OffsetDropTimeID           *string          `json:"offset_drop_time_id,omitempty" gorm:"type:varchar(36)"`
	Quantity                   *int             `json:"quantity,omitempty" gorm:"type:smallint"`
	Percent                    *float64         `json:"percent,omitempty"`
	CurrencyID                 *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" type:"uuid"`
	Amount                     *float64         `json:"amount,omitempty"`
	Description                *string          `json:"description,omitempty" gorm:"type:text"`
}
