package model

import (
	"log"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SqNdcCacheOrder | super class
type SqNdcCacheOrder struct {
	basic.Base
	basic.DataOwner
	SessionID                *uuid.UUID             `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                      // session ID
	OrderID                  *string                `json:"order_id,omitempty" gorm:"type:varchar(60)" default:"SQ_6V2FUV"`        // OrderID
	BookingID                *string                `json:"booking_id,omitempty" gorm:"type:varchar(60)" default:"6V2FUV"`         // BookingID
	BaseAmount               *float64               `json:"base_amount,omitempty" gorm:"type:decimal(22,2)"`                       // price information
	TotalAmount              *float64               `json:"total_amount,omitempty" gorm:"type:decimal(22,2)"`                      // price information
	TotalTaxAmount           *float64               `json:"total_tax_amount,omitempty" gorm:"type:decimal(22,2)"`                  // price information
	PenaltyAmount            *float64               `json:"penalty_amount,omitempty" gorm:"type:decimal(22,2)"`                    // PenaltyAmount when success void of an order
	ExpectedRefundAmount     *float64               `json:"expected_refund_amount,omitempty" gorm:"type:decimal(22,2)"`            // ExpectedRefundAmount when success void of an order
	PaymentTimeLimitDateTime *string                `json:"payment_time_limit_date_time,omitempty" gorm:"type:timestamptz"`        // PaymentTimeLimitDateTime
	OrderItem                *[]SqNdcCacheOrderItem `json:"order_item,omitempty" gorm:"foreignKey:SessionID;references:SessionID"` // OrderItem. Must be associated with SessionID, as it may mix session data when using OrderID in modify flight.
	Rules                    *basic.CustomDataTypes `json:"rules,omitempty"`                                                       // Rules meta data e.g : ATC_EXCHANGE_ELIGIBILITY / ATC_REFUND_ELIGIBILITY / VOID_ELIGIBILITY
}

// SqNdcCacheOrderItem | child class
type SqNdcCacheOrderItem struct {
	basic.Base
	basic.DataOwner
	SessionID                       *uuid.UUID                    `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                               // session ID
	OrderRefID                      *string                       `json:"order_id,omitempty" gorm:"type:varchar(60)" default:"SQ_6V2FUV"`                 // OrderRefID
	OrderItemID                     *string                       `json:"order_item_id,omitempty" gorm:"type:varchar(60)" default:"SQ_6V2FUV_AIR-1"`      // OrderItemID
	BaseAmount                      *float64                      `json:"base_amount,omitempty" gorm:"type:decimal(22,2)"`                                // price information
	TotalTaxAmount                  *float64                      `json:"total_tax_amount,omitempty" gorm:"type:decimal(22,2)"`                           // price information
	TotalAmount                     *float64                      `json:"total_amount,omitempty" gorm:"type:decimal(22,2)"`                               // price information
	PriceGuaranteeTimeLimitDateTime *string                       `json:"price_guarantee_time_limit_date_time,omitempty" gorm:"type:timestamptz"`         // PriceGuaranteeTimeLimitDateTime
	FareDetail                      *[]SqNdcCacheOrderFareDetail  `json:"fare_detail,omitempty" gorm:"foreignKey:OrderItemRefID;references:OrderItemID"`  // FareDetail
	Service                         *[]SqNdcCacheOrderService     `json:"service,omitempty" gorm:"foreignKey:OrderItemRefID;references:OrderItemID"`      // Service
	PaymentInfo                     *[]SqNdcCacheOrderPaymentInfo `json:"payment_info,omitempty" gorm:"foreignKey:OrderItemRefID;references:OrderItemID"` // PaymentInfo
}

// SqNdcCacheOrderFareDetail
type SqNdcCacheOrderFareDetail struct {
	basic.Base
	basic.DataOwner
	SessionID      *uuid.UUID                      `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                              // session ID
	OrderItemRefID *string                         `json:"order_item_ref_id,omitempty" gorm:"type:varchar(60)" default:"SQ_6V2FUV_AIR-1"` // OrderItemID
	PassengerRefs  *string                         `json:"passenger_refs,omitempty" gorm:"type:varchar(60)" default:"PAX1"`               // PassengerRefs
	CurrencyCode   *string                         `json:"currency_code,omitempty" gorm:"type:varchar(60)" default:"SGD"`                 // CurrencyCode
	BaseAmount     *float64                        `json:"base_amount,omitempty" gorm:"type:decimal(22,2)"`                               // price information
	TotalTaxAmount *float64                        `json:"total_tax_amount,omitempty" gorm:"type:decimal(22,2)"`                          // price information
	TotalAmount    *float64                        `json:"total_amount,omitempty" gorm:"type:decimal(22,2)"`                              // price information
	Remarks        *string                         `json:"remarks,omitempty"`                                                             // multiple value with comas || Remarks
	FareComponent  *[]SqNdcCacheOrderFareComponent `json:"fare_component,omitempty" gorm:"foreignKey:FareDetailID;references:ID"`         // FareComponent
	Taxes          *[]SqNdcCacheOrderTaxes         `json:"taxes,omitempty" gorm:"foreignKey:FareDetailID;references:ID"`                  // FareComponent
}

// SqNdcCacheOrderTaxes
type SqNdcCacheOrderTaxes struct {
	basic.Base
	basic.DataOwner
	SessionID    *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`     // session ID
	FareDetailID *uuid.UUID `json:"fare_detail_id,omitempty" gorm:"type:varchar(36)"`     // FareDetailID
	Code         *string    `json:"code,omitempty" gorm:"type:varchar(60)" default:"OP"`  // Code
	Amount       *float64   `json:"total_tax_amount,omitempty" gorm:"type:decimal(22,2)"` // price information
}

// SqNdcCacheOrderFareDetail
type SqNdcCacheOrderFareComponent struct {
	basic.Base
	basic.DataOwner
	SessionID     *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                    // session ID
	FareDetailID  *uuid.UUID `json:"fare_detail_id,omitempty" gorm:"type:varchar(36)"`                    // FareDetailID
	SegmentRefs   *string    `json:"segment_refs,omitempty" gorm:"type:varchar(60)" default:"SEG2"`       // SegmentRefs
	FareBasisCode *string    `json:"fare_basis_code,omitempty" gorm:"type:varchar(60)" default:"N16IAOL"` // FareBasisCode
	CabinTypeCode *string    `json:"cabin_type_code,omitempty" gorm:"type:varchar(10)" default:"M"`       // CabinTypeCode
	CabinTypeName *string    `json:"cabin_type_name,omitempty" gorm:"type:varchar(256)" default:"ECO"`    // CabinTypeName
	FareRules     *string    `json:"fare_rules,omitempty" gorm:"type:text"`
}

// SqNdcCacheOrderService
type SqNdcCacheOrderService struct {
	basic.Base
	basic.DataOwner
	SessionID              *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                                 // session ID
	OrderItemRefID         *string    `json:"order_item_ref_id,omitempty" gorm:"type:varchar(60)" default:"SQ_6V2FUV_AIR-1"`    // OrderItemID
	ServiceID              *string    `json:"service_id,omitempty" gorm:"type:varchar(60)" default:"SQ_6V2FUV_AIR"`             // ServiceID
	StatusCode             *string    `json:"status_code,omitempty" gorm:"type:varchar(60)" default:"HK"`                       // StatusCode
	PaxRefID               *string    `json:"pax_ref_id,omitempty" gorm:"type:varchar(60)" default:"PAX1"`                      // PaxRefID
	ServiceDefinitionRefID *string    `json:"service_definition_ref_id,omitempty" gorm:"type:varchar(60)" default:"BAGALLOW_1"` // ServiceDefinitionRefID
	PaxSegmentRefID        *string    `json:"pax_segment_ref_id,omitempty" gorm:"type:varchar(60)" default:"SEG1"`              // PaxSegmentRefID
	DatedOperatingLegRefID *string    `json:"dated_operating_leg_ref_id,omitempty" gorm:"type:varchar(60)" default:"SEG1"`      // PaxSegmentRefID
	SeatRow                *string    `json:"seat_row,omitempty" gorm:"type:varchar(60)" default:"50"`                          // SeatRow
	SeatColumn             *string    `json:"seat_column,omitempty" gorm:"type:varchar(60)" default:"A"`                        // SeatColumn
	SeatProfileRefID       *string    `json:"seat_profile_ref_id,omitempty" gorm:"type:varchar(60)" default:"RQST_1"`           // SeatProfileRefID
}

// SqNdcCacheOrderPaymentInfo
type SqNdcCacheOrderPaymentInfo struct {
	basic.Base
	basic.DataOwner
	SessionID      *uuid.UUID `json:"session_id" gorm:"type:varchar(36)" format:"uuid"`                              // session ID
	OrderItemRefID *string    `json:"order_item_ref_id,omitempty" gorm:"type:varchar(60)" default:"SQ_6V2FUV_AIR-1"` // OrderItemID
	Amount         *float64   `json:"amount,omitempty" gorm:"type:decimal(22,2)"`                                    // Amount
	TypeCode       *string    `json:"type_code,omitempty" gorm:"type:varchar(60)" default:"CASH"`                    // TypeCode
}

// SqNdcCacheOrderRules
type SqNdcCacheOrderRules struct {
	RuleID  *string `json:"rule_id,omitempty"`
	Status  *bool   `json:"status,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
}

// JsonOrderFareRules
type JsonOrderFareRules struct {
	Type                *string  `json:"type,omitempty" gorm:"type:varchar(36)" default:"Cancel"`   // enumable: Cancel, Change
	CurrencyAmountValue *float64 `json:"currency_amount_value,omitempty" gorm:"type:decimal(22,2)"` // CurrencyAmountValue
	AmountApplication   *string  `json:"amount_application" gorm:"type:varchar(60)" default:"ADM"`  // AmountApplication
}

// GetExistingOrder retrieves existing order data
func (data *SqNdcCacheOrder) GetExistingOrder(tx *gorm.DB, sessionID *uuid.UUID) {
	// Define function to preload order data
	orderItem := func(tx *gorm.DB) {
		tx.Preload("OrderItem.FareDetail.FareComponent").
			Preload("OrderItem.FareDetail.Taxes").
			Preload("OrderItem.Service").
			Preload("OrderItem.PaymentInfo").
			Take(&data)
	}

	// Try to get order data by session ID
	tx.Where("session_id = ?", sessionID).First(&data)
	if data.OrderID == nil {
		log.Printf("Current session (%s) does not have an Order history, trying to use the previous session", lib.ForceUUID(sessionID))

		// If not found, try to get from previous session
		flightSource := FlightCachingSources{}
		tx.Take(&flightSource, "session_id = ?", sessionID)
		if lib.ForceBool(flightSource.IsModifiedBooking) && flightSource.OldSessionID != nil {
			tx.Where("session_id = ?", flightSource.OldSessionID).First(&data)
		}
	}

	// Preload order data if found
	if data.OrderID != nil {
		orderItem(tx)
	}
}

// GetTotalAmount
func (s *SqNdcCacheOrderItem) GetTotalAmount(tx *gorm.DB, sessionID *uuid.UUID) (total string) {
	tx.Model(&s).Select("SUM(total_amount)").Where("session_id = ?", sessionID).Row().Scan(&total)
	return
}
