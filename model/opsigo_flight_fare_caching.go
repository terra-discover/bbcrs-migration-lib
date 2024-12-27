package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type OpsigoFlightFareCaching struct {
	basic.Base
	basic.DataOwner
	SessionID                *uuid.UUID `json:"session_id" gorm:"type:varchar(36)"`
	FareHash                 *uuid.UUID `json:"fare_hash" gorm:"type:varchar(36)"`         //random hash that will be used as shown "FareID"
	RouteTrxID               *string    `json:"route_trx_id" gorm:"type:varchar(255)"`     // the Opsigo Route TrxID related to this fare
	TrxID                    *string    `json:"trx_id,omitempty" gorm:"type:varchar(255)"` // unique 3rd party flight fare id
	FareName                 *string    `json:"fare_name,omitempty" gorm:"type:varchar(100)"`
	Price                    *float64   `json:"price,omitempty"`
	BagageMaxKg              *float64   `json:"bagage_max_kg,omitempty"`
	CurrencyCode             *string    `json:"currency_code,omitempty" gorm:"type:varchar(60)"`
	CurrencyName             *string    `json:"currency_name,omitempty" gorm:"type:varchar(60)"`
	SeatSelectionAllowed     *bool      `json:"seat_selection_allowed,omitempty"`
	SeatSelectionFee         *int       `json:"seat_selection_fee,omitempty"`
	KrisFlyerMilesPercentage *float64   `json:"kris_flyer_miles_percentage,omitempty"`
	UpgradeMilesAllowed      *bool      `json:"upgrade_miles_allowed,omitempty"`
	CancelationAllowed       *bool      `json:"cancelation_allowed,omitempty"`
	CancelationRefundable    *bool      `json:"cancelation_refundable,omitempty"`
	CancelationFee           *int       `json:"cancelation_fee,omitempty"`
	BookingChangeFeeBefore   *int       `json:"booking_change_fee_before,omitempty"`
	BookingChangeFeeAfter    *int       `json:"booking_change_fee_after,omitempty"`
	ChangeBookingBefore      *string    `json:"change_booking_before,omitempty" gorm:"type:varchar(60)"`
	ChangeBookingAfter       *string    `json:"change_booking_after,omitempty" gorm:"type:varchar(60)"`
	TripTypeCode             *string    `json:"trip_type_code,omitempty" gorm:"type:varchar(60)"`
	TripTypeName             *string    `json:"trip_type_name,omitempty" gorm:"type:varchar(60)"`
	NoShowFee                *int       `json:"no_show_fee,omitempty"`
	AvailableSeats           *int       `json:"available_seats,omitempty"`
}
