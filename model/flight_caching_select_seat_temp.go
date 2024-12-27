package model

// TODO: Since we are adopting a new flow, where the process of managing bookings/seats & addon selection uses a new session_id,
// this model FlightCachingSelectSeatTemp might not be used in the future (deprecated).

// FlightCachingSelectSeatTemp struct FlightCachingSelectSeatTemp
// type FlightCachingSelectSeatTemp struct {
// 	Base
// 	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_select_seat_temp_session_id;not null"`
// 	FlightCachingSelectSeatTempAPI
// 	SelectedRoute *FlightCachingRoutes `json:"selected_route,omitempty" gorm:"foreignKey:RouteID;references:ID"`
// 	SelectedFare  *FlightCachingFares  `json:"selected_fare,omitempty" gorm:"foreignKey:FareID;references:ID"`
// 	Person        *Person              `json:"person,omitempty" gorm:"foreignKey:PersonID;references:ID"`
// 	FlightSegment *FlightSegment       `json:"flight_segment,omitempty" gorm:"foreignKey:FlightSegmentID;references:ID"`
// 	IsNew         *bool                `json:"is_new,omitempty"` // untuk membedakan seat lama dan seat baru
// }

// // FlightCachingSelectSeatTempAPI FlightCachingSelectSeatTemp API
// type FlightCachingSelectSeatTempAPI struct {
// 	PersonID         *uuid.UUID `json:"person_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"` // PersonID
// 	RouteID          *uuid.UUID `json:"route_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"`  // leave to trace generic cache
// 	FareID           *uuid.UUID `json:"fare_id,omitempty" gorm:"varchar(36);not null" swaggertype:"string" format:"uuid"`   // leave to trace generic cache
// 	FlightSegmentID  *uuid.UUID `json:"flight_segment_id,omitempty" gorm:"varchar(36)" swaggertype:"string" format:"uuid"`  // summary records
// 	SeatCode         *string    `json:"seat_code,omitempty" gorm:"type:varchar(72);not null"`                               // summary records
// 	SeatColumn       *string    `json:"seat_column,omitempty" gorm:"type:varchar(1);not null"`                              // summary records
// 	SeatRow          *string    `json:"seat_row,omitempty" gorm:"type:varchar(3);not null"`                                 // summary records
// 	SeatNumber       *string    `json:"seat_number,omitempty" gorm:"type:varchar(4);not null"`                              // summary records
// 	CurrencyCode     *string    `json:"currency_code,omitempty" gorm:"type:varchar(36)"`
// 	Price            *float64   `json:"price,omitempty" gorm:"not null" validate:"required" example:"518000"` // summary records
// 	PriceAfterMarkup *float64   `json:"price_after_markup,omitempty"`                                         // will be filled by pricemodifier package
// 	IsPaid           *bool      `json:"is_paid,omitempty"`                                                    // IsPaid
// }
