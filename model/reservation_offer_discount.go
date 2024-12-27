package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationOfferDiscount struct {
	basic.Base
	basic.DataOwner
	ReservationOfferDiscountAPI
}

type ReservationOfferDiscountAPI struct {
	ReservationOfferID                      *uuid.UUID `json:"reservation_offer_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	OfferDiscountID                         *uuid.UUID `json:"offer_discount_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	DiscountCode                            *string    `json:"discount_code,omitempty" gorm:"type:varchar(36)"`
	DiscountName                            *string    `json:"discount_name,omitempty" gorm:"type:varchar(64)"`
	ChargeTypeID                            *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	AgeQualifyingGroupID                    *uuid.UUID `json:"age_qualifying_group_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ApplicationOrder                        *int       `json:"application_order,omitempty" gorm:"type:int"`
	NumberOfNightsRequired                  *int       `json:"number_of_nights_required,omitempty" gorm:"type:smallint"`
	NumberOfNightsDiscounted                *int       `json:"number_of_nights_discounted,omitempty" gorm:"type:smallint"`
	DiscountPattern                         *string    `json:"discount_pattern,omitempty" gorm:"type:varchar(36)"`
	DiscountPercent                         *float64   `json:"discount_percent,omitempty"`
	CurrencyID                              *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	DiscountAmount                          *float64   `json:"discount_amount,omitempty"`
	IsDiscountTaxInclusive                  *bool      `json:"is_discount_tax_inclusive,omitempty"`
	IsDiscountServiceFeeInclusive           *bool      `json:"is_discount_service_fee_inclusive,omitempty"`
	IsDiscountFeeInclusive                  *bool      `json:"is_discount_fee_inclusive,omitempty"`
	IsDiscountPreselectedComponentInclusive *bool      `json:"is_discount_preselected_component_inclusive,omitempty"`
	IsDiscountSellableProductInclusive      *bool      `json:"is_discount_sellable_product_inclusive,omitempty"`
	DiscountReason                          *string    `json:"discount_reason,omitempty" gorm:"type:text"`
}
