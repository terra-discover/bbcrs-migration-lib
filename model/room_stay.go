package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomStay Room Stay
type RoomStay struct {
	basic.Base
	basic.DataOwner
	RoomStayAPI
	Hotel           *Hotel           `json:"hotel,omitempty"`
	RoomStayGuest   []RoomStayGuest  `json:"room_stay_guest,omitempty"`
	RoomStayPayment *RoomStayPayment `json:"room_stay_payment,omitempty"`
	RoomRate        []RoomRate       `json:"room_rate,omitempty"`
}

// RoomStayAPI Room Stay API
type RoomStayAPI struct {
	HotelID                          *uuid.UUID       `json:"hotel_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                             // Hotel ID
	RoomStayStatus                   *int             `json:"room_stay_status,omitempty" gorm:"type:smallint"`                                                           // Room Stay Status
	Checkin                          *strfmt.DateTime `json:"checkin,omitempty" gorm:"type:timestamptz" swaggertype:"string" format:"date-time"`                         // Checkin
	Checkout                         *strfmt.DateTime `json:"checkout,omitempty" gorm:"type:timestamptz" swaggertype:"string" format:"date-time"`                        // Checkout
	BaseCurrencyID                   *uuid.UUID       `json:"base_price_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`               // Base Price Currency ID
	BaseAmount                       *float64         `json:"base_price_amount,omitempty" gorm:"type:decimal(19,4)"`                                                     // Base Price Amount
	EquivalentPriceCurrencyID        *uuid.UUID       `json:"equivalent_price_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`         // Equivalent Price Currency ID
	EquivalentPriceAmount            *float64         `json:"equivalent_price_amount,omitempty" gorm:"type:decimal(19,4)"`                                               // Equivalent Price Amount
	TotalTaxCurrencyID               *uuid.UUID       `json:"total_tax_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                // Total Tax Currency ID
	TotalTaxAmount                   *float64         `json:"total_tax_amount,omitempty" gorm:"type:decimal(19,4)"`                                                      // Total Tax Amount
	TotalFeeCurrencyID               *uuid.UUID       `json:"total_fee_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                // Total Fee Currency ID
	TotalFeeAmount                   *float64         `json:"total_fee_amount,omitempty" gorm:"type:decimal(19,4)"`                                                      // Total Fee Amount
	TotalPriceCurrencyID             *uuid.UUID       `json:"total_price_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // Total Price Currency ID
	TotalPriceAmount                 *float64         `json:"total_price_amount,omitempty" gorm:"type:decimal(19,4)"`                                                    // Total Price Amount
	CurrencyID                       *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                          // Currency ID
	TotalAmountBeforeTax             *float64         `json:"total_amount_before_tax,omitempty" gorm:"type:decimal(19,4)"`                                               // Total Amount Before Tax
	TotalAmountAfterTax              *float64         `json:"total_amount_after_tax,omitempty" gorm:"type:decimal(19,4)"`                                                // Total Amount After Tax
	CommissionPercent                *float64         `json:"commission_percent,omitempty" gorm:"type:decimal(19,4)"`                                                    // Commission Percent
	CommissionCurrencyID             *uuid.UUID       `json:"commission_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`               // Commission Currency ID
	CommissionAmount                 *float64         `json:"commission_amount,omitempty" gorm:"type:decimal(19,4)"`                                                     // Commission Amount
	BasePenaltyCurrencyID            *uuid.UUID       `json:"base_penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`             // BasePenal ty Currency ID
	BasePenaltyAmount                *float64         `json:"base_penalty_amount,omitempty" gorm:"type:decimal(19,4)"`                                                   // BasePenal ty Amount
	EquivalentPenaltyCurrencyID      *uuid.UUID       `json:"equivalent_penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`       // Equivalent Penalty Currency ID
	EquivalentPenaltyAmount          *float64         `json:"equivalent_penalty_amount,omitempty" gorm:"type:decimal(19,4)"`                                             // Equivalent Penalty Amount
	TotalPenaltyTaxCurrencyID        *uuid.UUID       `json:"total_penalty_tax_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`        // Total PenaltyTax Currency ID
	TotalPenaltyTaxAmount            *float64         `json:"total_penalty_tax_amount,omitempty" gorm:"type:decimal(19,4)"`                                              // Total PenaltyTax Amount
	TotalPenaltyCurrencyID           *uuid.UUID       `json:"total_penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Total Penalty Currency ID
	TotalPenaltyAmount               *float64         `json:"total_penalty_amount,omitempty" gorm:"type:decimal(19,4)"`                                                  // Total Penalty Amount
	PenaltyCurrencyID                *uuid.UUID       `json:"penalty_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                  // Penalty Currency ID
	TotalPenaltyAmountBeforeTax      *float64         `json:"total_penalty_amount_before_tax,omitempty" gorm:"type:decimal(19,4)"`                                       // Total Penalty Amount Before Tax
	TotalPenaltyAmountAfterTax       *float64         `json:"total_penalty_amount_after_tax,omitempty" gorm:"type:decimal(19,4)"`                                        // Total Penalty Amount After Tax
	CurrencyConversionFromCurrencyID *uuid.UUID       `json:"currency_conversion_from_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // Currency Conversion From Currency ID
	CurrencyConversionToCurrencyID   *uuid.UUID       `json:"currency_conversion_to_currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`   // Currency Conversion To Currency ID
	CurrencyConversionMultiplyRate   *float64         `json:"currency_conversion_multiply_rate,omitempty" gorm:"type:decimal(19,4)"`                                     // Currency Conversion Multiply Rate
	HotelCancellationDeadline        *strfmt.DateTime `json:"hotel_cancellation_deadline,omitempty" gorm:"type:timestamptz" swaggertype:"string" format:"date-time"`     // Hotel Cancellation Deadline
	AgentCancellationDeadline        *strfmt.DateTime `json:"agent_cancellation_deadline,omitempty" gorm:"type:timestamptz" swaggertype:"string" format:"date-time"`     // Agent Cancellation Deadline
	CancelledAt                      *strfmt.DateTime `json:"cancelled_at,omitempty" gorm:"type:timestamptz" swaggertype:"string" format:"date-time"`                    // Cancelled At
	CancellationReason               *string          `json:"cancellation_reason,omitempty" gorm:"type:text"`                                                            // Cancellation Reason
	TaxCategoryID                    *uuid.UUID       `json:"tax_category_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                      // Tax Category ID
	HotelSupplierID                  *uuid.UUID       `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                    // Hotel Supplier ID
	SupplierCode                     *string          `json:"supplier_code,omitempty" gorm:"type:varchar(16)"`                                                           // Supplier Code
	SupplierSubCode                  *string          `json:"supplier_sub_code,omitempty" gorm:"type:varchar(256)"`                                                      // Supplier Sub Code
	SupplierSubHotelCode             *string          `json:"supplier_sub_hotel_code,omitempty" gorm:"type:varchar(64)"`                                                 // Supplier Sub Hotel Code
	HotelAcknowledgmentStatus        *int             `json:"hotel_acknowledgment_status,omitempty" gorm:"type:smallint"`                                                // Hotel Acknowledgment Status
	HotelAcknowledgmentComment       *string          `json:"hotel_acknowledgment_comment,omitempty" gorm:"type:text"`                                                   // Hotel Acknowledgment Comment
	HotelAcknowledgmentBy            *string          `json:"hotel_acknowledgment_by,omitempty" gorm:"type:varchar(64)"`                                                 // Hotel Acknowledgment By
	Comment                          *string          `json:"comment,omitempty" gorm:"type:text"`                                                                        // Comment
	ConsumerViewableComment          *string          `json:"consumer_viewable_comment,omitempty" gorm:"type:text"`                                                      // Consumer Viewable Comment

}
