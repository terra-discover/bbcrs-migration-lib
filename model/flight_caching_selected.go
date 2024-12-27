package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlightCachingSelected struct {
	basic.Base
	basic.DataOwner
	SessionID *uuid.UUID `json:"session_id,omitempty" gorm:"type:varchar(36);not null;index:idx_selected_airline_fare,unique"`
	FlightCachingSelectedAPI
	SelectedAirline     FlightCachingAirlines    `json:"selected_airline,omitempty" gorm:"foreignKey:AirlineID;references:ID"`
	SelectedFare        FlightCachingFares       `json:"selected_fare,omitempty" gorm:"foreignkey:FareID;references:ID"`
	SelectedFlightRoute []FlightCachingRoutes    `json:"selected_route,omitempty" gorm:"-"`
	SelectedTravelers   []FlightCachingTraveller `json:"-" gorm:"-"`
}

type FlightCachingSelectedAPI struct {
	AirlineID         *uuid.UUID       `json:"airline_id,omitempty" validate:"required" gorm:"type:varchar(36);not null;index:idx_selected_airline_fare,unique" swaggertype:"string" format:"uuid"` // Refer to FlightCachingAirlines.ID
	FareID            *uuid.UUID       `json:"fare_id,omitempty" validate:"required" gorm:"type:varchar(36);not null;index:idx_selected_airline_fare,unique" swaggertype:"string" format:"uuid"`    // Refer to FlightCachingFares.ID
	DepartureDate     *strfmt.DateTime `json:"departure_date,omitempty" swaggerignore:"true"`                                                                                                       // sorting field helper (help sorting cart by its date)
	ViolationReasonID *uuid.UUID       `json:"violation_reason_id" gorm:"type:varchar(36)"`
}

// GetFlightCachingSelected by query filter
func (data *FlightCachingSelected) GetFlightCachingSelected(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Preload(`SelectedAirline.FlightCachingRoutes`).Preload(`SelectedFare`).Take(&data)
}

// GetFlightCachingSelected by query filter
func (data *FlightCachingSelected) GetFlightCachingSelectedBySessionID(tx *gorm.DB, sessionID *uuid.UUID) (res []FlightCachingSelected) {
	tx.Preload("SelectedAirline").
		Preload("SelectedAirline.FlightCachingRoutes").
		Preload("SelectedFare.FlightCachingFareDetail.FlightTraveller").
		Where("session_id = ?", sessionID).
		Order(`sort ASC`).
		Find(&res)
	return
}
