package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// FlightCachingFares model
type FlightCachingFares struct {
	basic.Base
	basic.DataOwner
	SessionID                     *uuid.UUID                 `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_fares_session_id;not null"`
	FlightCachingCabinsID         *uuid.UUID                 `json:"flight_caching_cabins_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_fares_flight_caching_cabins_id;not null" format:"uuid"` // FlightCachingCabinsID
	TrxID                         *string                    `json:"trx_id,omitempty" gorm:"type:varchar(200)"`
	FareBasisCode                 *string                    `json:"fare_basis_code"`
	FareName                      *string                    `json:"fare_name,omitempty" gorm:"type:varchar(255)"`
	PriceAfterMarkup              *float64                   `json:"price_after_markup,omitempty"`
	CommissionClaimID             *uuid.UUID                 `json:"commission_claim_id,omitempty" gorm:"type:varchar(36)"`
	MarkupRateID                  *uuid.UUID                 `json:"markup_rate_id,omitempty" gorm:"type:varchar(36)"`
	ServiceFeeRateID              *uuid.UUID                 `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	TaxRateID                     *uuid.UUID                 `json:"tax_rate_id,omitempty" gorm:"type:varchar(36)"`
	ServiceFeeTaxRateID           *uuid.UUID                 `json:"service_fee_tax_rate_id,omitempty" gorm:"type:varchar(36)"`
	MDRProcessingFeeRateID        *uuid.UUID                 `json:"mdr_processing_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	RevalidateProcessingFeeRateID *uuid.UUID                 `json:"revalidate_processing_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	ReissueProcessingFeeRateID    *uuid.UUID                 `json:"reissue_processing_fee_rate_id,omitempty" gorm:"type:varchar(36)"`
	Price                         *float64                   `json:"price,omitempty"`              // Current or New price amount | bundling price, price after markup
	PrevPartnerPrice              *float64                   `json:"prev_partner_price,omitempty"` // Prevous partner price amount | already include tax | used for module modify flight before or after issued only
	PrevBasePrice                 *float64                   `json:"prev_base_price,omitempty"`    // Prevous base price amount
	PrevTotalTax                  *float64                   `json:"prev_total_tax,omitempty"`     // Prevous total tax amount
	PartnerPrice                  *float64                   `json:"partner_price,omitempty"`      // Current or New partner price amount
	BasePrice                     *float64                   `json:"base_price,omitempty"`         // Current or New base price amount
	TotalTax                      *float64                   `json:"total_tax,omitempty"`          // Current or New total tax amount
	PenaltyPrice                  *float64                   `json:"penalty_price,omitempty"`      // Represents the penalty price fee, if applicable | used for module modify flight after issued only
	PriceDifference               *float64                   `json:"price_difference,omitempty"`   // Difference between new and previous prices (exclude penalty price). Positive: Upgrade, Negative: Downgrade | used for module modify flight before or after issued only
	PriceAmountDue                *float64                   `json:"price_amount_due,omitempty"`   // Represents the amount due to pay for the current pricing |  used for module modify flight after issued only
	EquivalentFareAmount          *float64                   `json:"equivalent_fare_amount"`
	EquivalentFareCurrencyCode    *string                    `json:"equivalent_fare_currency_code,omitempty"`
	ComparePrice                  *ComparePrice              `json:"compare_schedules,omitempty" gorm:"-"` // price information to compare deviation price between fares (Only used when search flight view by schedule mechanism)
	BaggageMaxKg                  *float64                   `json:"baggage_max_kg,omitempty"`
	BaggageItems                  *basic.StringArray         `json:"baggage_items,omitempty"`
	CurrencyCode                  *string                    `json:"currency_code,omitempty" gorm:"type:varchar(36)"`
	CurrencyName                  *string                    `json:"currency_name,omitempty" gorm:"type:varchar(255)"`
	SeatSelectionAllowed          *bool                      `json:"seat_selection_allowed,omitempty"`
	SeatSelectionFee              *float64                   `json:"seat_selection_fee,omitempty"`
	SeatSelectionTypes            *string                    `json:"seat_selection_types,omitempty"`
	SeatSelectionDesc             *string                    `json:"seat_selection_desc,omitempty"`
	KrisFlyerMilesPercentage      *float64                   `json:"kris_flyer_miles_percentage,omitempty"`
	UpgradeMilesAllowed           *bool                      `json:"upgrade_miles_allowed,omitempty"`
	UpgradeMilesDesc              *string                    `json:"upgrade_miles_desc,omitempty"`
	CancelationAllowed            *bool                      `json:"cancelation_allowed,omitempty"`
	CancelationDesc               *string                    `json:"cancelation_desc,omitempty"`
	CancelationRefundable         *bool                      `json:"cancelation_refundable,omitempty"`
	CancelationFee                *int                       `json:"cancelation_fee,omitempty"`
	BookingChangeFeeBefore        *int                       `json:"booking_change_fee_before,omitempty"`
	BookingChangeFeeAfter         *int                       `json:"booking_change_fee_after,omitempty"`
	BookingChangeDesc             *string                    `json:"booking_change_desc,omitempty"`
	ChangeBookingBefore           *string                    `json:"change_booking_before,omitempty" gorm:"type:varchar(255)"`
	ChangeBookingAfter            *string                    `json:"change_booking_after,omitempty" gorm:"type:varchar(255)"`
	CancellationBefore            *string                    `json:"cancellation_before,omitempty" gorm:"type:varchar(255)"`
	CancellationAfter             *string                    `json:"cancellation_after,omitempty" gorm:"type:varchar(255)"`
	TripTypeCode                  *string                    `json:"trip_type_code,omitempty" gorm:"type:varchar(36)"`
	TripTypeName                  *string                    `json:"trip_type_name,omitempty" gorm:"type:varchar(255)"`
	NoShowFee                     *int                       `json:"no_show_fee,omitempty"`
	NoShowDesc                    *string                    `json:"no_show_desc,omitempty"`
	AvailableSeats                *int                       `json:"available_seats,omitempty"`
	Additional                    *string                    `json:"additional,omitempty"`
	VendorCode                    *string                    `json:"vendor_code,omitempty" gorm:"type:varchar(16)"`
	FareType                      *string                    `json:"fare_type,omitempty" gorm:"type:varchar(16)"`
	FareTariff                    *string                    `json:"fare_tariff,omitempty" gorm:"type:varchar(16)"`
	FareCode                      *string                    `json:"fare_code,omitempty" gorm:"type:varchar(16)"`
	FareRulesRemarks              *string                    `json:"fare_rules_remarks,omitempty" gorm:"type:varchar(191)"`
	ClassRef                      *string                    `json:"class_ref,omitempty" gorm:"type:varchar(36)"`
	ClassCode                     *string                    `json:"class_code,omitempty" gorm:"type:varchar(36)"`
	ClassName                     *string                    `json:"class_name,omitempty" gorm:"type:varchar(36)"`
	IsPrivateFare                 *bool                      `json:"is_private_fare,omitempty"`
	AvailableCabinTypes           *basic.CustomDataTypes     `json:"available_cabin_types,omitempty"`
	FareInfo                      *basic.CustomDataTypes     `json:"fare_info,omitempty"`
	FlightCachingFareDetail       *[]FlightCachingFareDetail `json:"fare_detail,omitempty" gorm:"foreignKey:FlightCachingFareID;references:ID"`
	FlightCachingCabins           *FlightCachingCabins       `json:"flight_caching_cabins,omitempty" gorm:"foreignKey:FlightCachingCabinsID"`
	PricingUsed                                              // Only show in development
}

// Only show in development
type PricingUsed struct {
	UsedMarkupRate        *MarkupRate     `json:"used_markup_rate,omitempty" gorm:"foreignKey:MarkupRateID"`
	UsedServiceFeeRate    *ServiceFeeRate `json:"used_service_fee_rate,omitempty" gorm:"foreignKey:ServiceFeeRateID"`
	UsedTaxRate           *TaxRate        `json:"used_tax_rate,omitempty" gorm:"foreignKey:TaxRateID"`
	UsedServiceFeeTaxRate *TaxRate        `json:"used_service_fee_tax_rate,omitempty" gorm:"foreignKey:ServiceFeeTaxRateID"`
}

// ComparePrice
type ComparePrice struct {
	IsUpgrade         *bool    `json:"is_upgrade"`
	PriceAfterCompare *float64 `json:"price_after_compare"`
}
