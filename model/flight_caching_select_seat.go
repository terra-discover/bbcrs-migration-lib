package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FlightCachingSelectSeat | Do not delete this model, because it is used as a pivot table to continue transactions booking as information seat
type FlightCachingSelectSeat struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_select_seat_session_id;not null"`
	FlightCachingSelectSeatAPI
	SelectedRoute              *FlightCachingRoutes        `json:"selected_route,omitempty" gorm:"foreignKey:RouteID;references:ID"`
	SelectedFare               *FlightCachingFares         `json:"selected_fare,omitempty" gorm:"foreignKey:FareID;references:ID"`
	Person                     *Person                     `json:"person,omitempty" gorm:"foreignKey:PersonID;references:ID"`
	FlightSegment              *FlightSegment              `json:"flight_segment,omitempty" gorm:"foreignKey:FlightSegmentID;references:ID"`
	ProposalProduct            *ProposalProduct            `json:"proposal_product,omitempty" gorm:"foreignKey:ProposalProductID;references:ID"`
	ReservationProduct         *ReservationProduct         `json:"reservation_product,omitempty" gorm:"foreignKey:ReservationProductID;references:ID"`
	ProposalProductPurchase    *ProposalProductPurchase    `json:"proposal_product_purchase,omitempty" gorm:"foreignKey:ProposalProductPurchaseID;references:ID"`
	ReservationProductPurchase *ReservationProductPurchase `json:"reservation_product_purchase,omitempty" gorm:"foreignKey:ReservationProductPurchaseID;references:ID"`
	ProposalFlightSegment      *AirTravelerFlightSegment   `json:"proposal_flight_segment,omitempty" gorm:"foreignKey:ProposalFlightSegmentID;references:ID"`
	ReservationFlightSegment   *AirTravelerFlightSegment   `json:"reservation_flight_segment,omitempty" gorm:"foreignKey:ReservationFlightSegmentID;references:ID"`
	// IsNew                      *bool                       `json:"is_new,omitempty"` // untuk membedakan seat lama dan seat baru [!] DEPRECATED
}

// FlightCachingSelectSeatAPI
type FlightCachingSelectSeatAPI struct {
	PersonID                     *uuid.UUID `json:"person_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"`              // PersonID
	RouteID                      *uuid.UUID `json:"route_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"`               // leave to trace generic cache
	FareID                       *uuid.UUID `json:"fare_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"`                // leave to trace generic cache
	FlightSegmentID              *uuid.UUID `json:"flight_segment_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"`      // summary records
	ProposalProductID            *uuid.UUID `json:"proposal_product_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"`             // leave to trace generic cache to propher spesifications
	ReservationProductID         *uuid.UUID `json:"reservation_product_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"`          // leave to trace generic cache to propher spesifications
	ProposalProductPurchaseID    *uuid.UUID `json:"proposal_product_purchase_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"`    // leave to trace generic cache to propher spesifications
	ReservationProductPurchaseID *uuid.UUID `json:"reservation_product_purchase_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"` // leave to trace generic cache to propher spesifications
	ProposalFlightSegmentID      *uuid.UUID `json:"proposal_flight_segment_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"`      // leave to trace generic cache to propher spesifications
	ReservationFlightSegmentID   *uuid.UUID `json:"reservation_flight_segment_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"`   // leave to trace generic cache to propher spesifications
	SeatCode                     *string    `json:"seat_code,omitempty" gorm:"type:text;not null"`                                                   // summary records
	SeatColumn                   *string    `json:"seat_column,omitempty" gorm:"type:varchar(1);not null"`                                           // summary records
	SeatRow                      *string    `json:"seat_row,omitempty" gorm:"type:varchar(3);not null"`                                              // summary records
	SeatNumber                   *string    `json:"seat_number,omitempty" gorm:"type:varchar(4);not null"`                                           // summary records
	CurrencyCode                 *string    `json:"currency_code,omitempty" gorm:"type:varchar(36)"`
	Price                        *float64   `json:"price,omitempty" gorm:"not null" validate:"required" example:"518000"` // summary records
	PriceAfterMarkup             *float64   `json:"price_after_markup,omitempty"`                                         // will be filled by pricemodifier package
	IsCancelled                  *bool      `json:"is_cancelled,omitempty"`                                               // IsCancelled
	IsPaid                       *bool      `json:"is_paid,omitempty"`                                                    // IsPaid
}

// BeforeCreate Data
func (b *FlightCachingSelectSeat) BeforeCreate(tx *gorm.DB) error {
	b.Base.BeforeCreate(tx)
	if nil == b.IsCancelled {
		b.IsCancelled = lib.Boolptr(false)
	}
	return nil
}

// GetFlightCachingSelectSeat by query filter
func (data *FlightCachingSelectSeat) GetFlightCachingSelectSeat(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
}
