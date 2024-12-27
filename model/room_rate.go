package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomRate Room Rate
type RoomRate struct {
	basic.Base
	basic.DataOwner
	RoomRateAPI
	ReservationRoomType *ReservationRoomType `json:"reservation_room_type,omitempty" gorm:"foreignKey:RoomTypeID;references:RoomTypeID;"`
	ProposalRoomType    *ProposalRoomType    `json:"proposal_room_type,omitempty" gorm:"foreignKey:RoomTypeID;references:RoomTypeID;"`
}

// RoomRateAPI Room Rate Api
type RoomRateAPI struct {
	RoomStayID                       *uuid.UUID       `json:"room_stay_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	RoomTypeID                       *uuid.UUID       `json:"room_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	RatePlanID                       *uuid.UUID       `json:"rate_plan_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	OfferID                          *uuid.UUID       `json:"offer_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	NumberOfRooms                    *int             `json:"number_of_rooms,omitempty" gorm:"type:smallint" swaggertype:"number" format:"int32" example:"10"`
	StartDate                        *strfmt.DateTime `json:"start_date,omitempty" gorm:"type:timestamptz" swaggertype:"string" format:"date-time"`
	EndDate                          *strfmt.DateTime `json:"end_date,omitempty" gorm:"type:timestamptz" swaggertype:"string" format:"date-time"`
	BaseCurrencyID                   *uuid.UUID       `json:"base_price_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	BaseAmount                       *float64         `json:"base_price_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	EquivalentPriceCurrencyID        *uuid.UUID       `json:"equivalent_price_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	EquivalentPriceAmount            *float64         `json:"equivalent_price_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	TotalTaxCurrencyID               *uuid.UUID       `json:"total_tax_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	TotalTaxAmount                   *float64         `json:"total_tax_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	TotalFeeCurrencyID               *uuid.UUID       `json:"total_fee_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	TotalFeeAmount                   *float64         `json:"total_fee_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	TotalPriceCurrencyID             *uuid.UUID       `json:"total_price_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	TotalPriceAmount                 *float64         `json:"total_price_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	CurrencyID                       *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	TotalAmountBeforeTax             *float64         `json:"total_amount_before_tax,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	TotalAmountAfterTax              *float64         `json:"total_amount_after_tax,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	CommissionPercent                *float64         `json:"commission_percent,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	CommissionCurrencyID             *uuid.UUID       `json:"commission_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	CommissionAmount                 *float64         `json:"commission_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	BasePenaltyCurrencyID            *uuid.UUID       `json:"base_penalty_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	BasePenaltyAmount                *float64         `json:"base_penalty_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	EquivalentPenaltyCurrencyID      *uuid.UUID       `json:"equivalent_penalty_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	EquivalentPenaltyAmount          *float64         `json:"equivalent_penalty_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	TotalPenaltyTaxCurrencyID        *uuid.UUID       `json:"total_penalty_tax_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	TotalPenaltyTaxAmount            *float64         `json:"total_penalty_tax_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	TotalPenaltyCurrencyID           *uuid.UUID       `json:"total_penalty_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	TotalPenaltyAmount               *float64         `json:"total_penalty_amount,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	PenaltyCurrencyID                *uuid.UUID       `json:"penalty_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	TotalPenaltyAmountBeforeTax      *float64         `json:"total_penalty_amount_before_tax,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	TotalPenaltyAmountAfterTax       *float64         `json:"total_penalty_amount_after_tax,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
	CurrencyConversionFromCurrencyID *uuid.UUID       `json:"currency_conversion_from_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	CurrencyConversionToCurrencyID   *uuid.UUID       `json:"currency_conversion_to_currency_id,omitempty" gorm:"type:varchar(36);" swaggertype:"string" format:"uuid"`
	CurrencyConversionMultiplyRate   *float64         `json:"currency_conversion_multiply_rate,omitempty" gorm:"type:decimal(19,4)" swaggertype:"number" format:"float"`
}
