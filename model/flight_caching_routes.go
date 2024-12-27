package model

import (
	"fmt"
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FlightCachingRoutes model
type FlightCachingRoutes struct {
	basic.Base
	basic.DataOwner
	SessionID                *uuid.UUID             `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_routes_session_id;not null"`
	FlightCachingAirlinesID  *uuid.UUID             `json:"flight_caching_airlines_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_routes_flight_caching_airlines_id;not null" format:"uuid"` // FlightCachingID
	TrxID                    *string                `json:"trx_id,omitempty"`                                                                                                                               //unique 3rd party flight routes id
	OldTrxID                 *string                `json:"old_trx_id,omitempty" gorm:"type:varchar(50)"`
	FareBasisCode            *string                `json:"fare_basis_code,omitempty" gorm:"type:varchar(36)"`
	FareClass                *string                `json:"fare_class,omitempty"`
	FareCode                 *string                `json:"fare_code,omitempty"`
	FareName                 *string                `json:"fare_name,omitempty" gorm:"type:varchar(36)"`
	AirlineCode              *string                `json:"airline_code,omitempty" gorm:"type:varchar(36)"`
	AirlineName              *string                `json:"airline_name,omitempty" gorm:"type:varchar(255)"`
	SourceCode               *string                `json:"source_code,omitempty" gorm:"type:varchar(36);not null"`
	SourceNumber             *string                `json:"source_number,omitempty" gorm:"type:varchar(255)"`
	AircraftCode             *string                `json:"aircraft_code,omitempty" gorm:"type:varchar(100)"`
	AircraftName             *string                `json:"aircraft_name,omitempty" gorm:"type:varchar(255)"`
	Index                    *int                   `json:"index,omitempty"`
	Mileage                  *int                   `json:"mileage,omitempty"`
	MileageUnitOfMeasureName *string                `json:"mileage_unit_of_measure_name,omitempty" gorm:"-"`
	PreviousArrivalDate      *strfmt.DateTime       `json:"previous_arrival_date,omitempty" format:"date-time" gorm:"type:timestamptz" default:"2022-06-05T10:10:00" swaggertype:"string"`
	DepartureDate            *strfmt.DateTime       `json:"departure_date,omitempty" format:"date-time" gorm:"type:timestamptz" default:"2022-06-05T10:10:00" swaggertype:"string"`
	DepartureGmtOffset       *float64               `json:"departure_gmt_offset,omitempty" example:"7.0"`
	DepartureAirportCode     *string                `json:"departure_airport_code,omitempty" gorm:"type:varchar(36)"`
	DepartureAirportName     *string                `json:"departure_airport_name,omitempty" gorm:"type:varchar(255)"`
	DepartureCityCode        *string                `json:"departure_city_code,omitempty" gorm:"type:varchar(36)"`
	DepartureCityName        *string                `json:"departure_city_name,omitempty" gorm:"type:varchar(255)"`
	DepartureCityID          *uuid.UUID             `json:"departure_city_id" gorm:"type:varchar(36)"`
	DepartureCountryID       *uuid.UUID             `json:"departure_country_id" gorm:"type:varchar(36)"`
	DepartureTerminal        *string                `json:"departure_terminal,omitempty" gorm:"type:varchar(255)"`
	ArrivalDate              *strfmt.DateTime       `json:"arrival_date,omitempty" format:"date-time" gorm:"type:timestamptz" default:"2022-06-05T10:10:00"  swaggertype:"string"`
	ArrivalGmtOffset         *float64               `json:"arrival_gmt_offset,omitempty" example:"8.0"`
	ArrivalAirportCode       *string                `json:"arrival_airport_code,omitempty" gorm:"type:varchar(36)"`
	ArrivalAirportName       *string                `json:"arrival_airport_name,omitempty" gorm:"type:varchar(255)"`
	ArrivalCityCode          *string                `json:"arrival_city_code,omitempty" gorm:"type:varchar(36)"`
	ArrivalCityName          *string                `json:"arrival_city_name,omitempty" gorm:"type:varchar(255)"`
	ArrivalCityID            *uuid.UUID             `json:"arrival_city_id" gorm:"type:varchar(36)"`
	ArrivalCountryID         *uuid.UUID             `json:"arrival_country_id" gorm:"type:varchar(36)"`
	ArrivalTerminal          *string                `json:"arrival_terminal,omitempty" gorm:"type:varchar(255)"`
	UnixTransitDuration      *int64                 `json:"unix_transit_duration"`
	TransitDuration          *string                `json:"transit_duration" gorm:"type:varchar(30)"`
	CabinTypeCode            *string                `json:"cabin_type_code,omitempty" gorm:"type:varchar(36)"`
	CabinTypeName            *string                `json:"cabin_type_name,omitempty" gorm:"type:varchar(255)"`
	Meal                     *string                `json:"meal,omitempty"`
	Baggage                  *float64               `json:"baggage,omitempty"`
	BaggageInfo              *string                `json:"baggage_info,omitempty"`
	FacilityBagage           *bool                  `json:"facility_bagage,omitempty"`
	Sequence                 *int                   `json:"sequence,omitempty"`
	FacilityWifi             *bool                  `json:"facility_wifi,omitempty"`
	FacilityMeal             *bool                  `json:"facility_meal,omitempty"`
	FacilitySeatSelection    *bool                  `json:"facility_seat_selection,omitempty"`
	FacilityProtection       *bool                  `json:"facility_protection,omitempty"`
	FacilityCharger          *bool                  `json:"facility_charger,omitempty"`
	StringDuration           *string                `json:"string_duration,omitempty"`
	UnixDuration             *int64                 `json:"unix_duration,omitempty"`
	Rules                    *basic.StringArray     `json:"rules,omitempty" gorm:"type:text"` // multiple value separates by comma
	FlightCachingAirlines    *FlightCachingAirlines `json:"airlines,omitempty" gorm:"foreignKey:FlightCachingAirlinesID;references:ID"`
}

// GetFlightCachingRoutes by query filter
func (data *FlightCachingRoutes) GetFlightCachingRoutes(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	if err := tx.Select(`id, session_id, flight_caching_airlines_id, trx_id, airline_code, source_code, source_number, aircraft_code, aircraft_name, mileage, previous_arrival_date, departure_date, departure_airport_code, departure_airport_name, departure_city_code, departure_city_name, departure_terminal, arrival_date, arrival_airport_code, arrival_airport_name, arrival_city_code, arrival_city_name, arrival_terminal, cabin_type_code, cabin_type_name, facility_bagage, facility_wifi, facility_meal, facility_seat_selection, facility_protection, facility_charger`).
		Where(qf, wf...).Take(&data); err.Error != nil {
		fmt.Println("err : ", err.Error)
	}
}

func (data *FlightCachingRoutes) GetFareBasisCode(tx *gorm.DB) *string {
	// SQ case
	if data.FareBasisCode != nil {
		return data.FareBasisCode
	}

	// OPSIGO must get them from Fare
	if data.TrxID != nil {
		selected := []FlightCachingSelected{}
		tx.Preload("SelectedAirline.FlightCachingRoutes").
			Preload("SelectedFare").
			Where("session_id = ?", data.SessionID).
			Find(&selected)

		var fareID *uuid.UUID
		for _, sel := range selected {
			if sel.SelectedAirline.FlightCachingRoutes != nil {
				for _, route := range *sel.SelectedAirline.FlightCachingRoutes {
					if route.ID.String() == data.ID.String() {
						fareID = sel.SelectedFare.ID
					}
				}
			}
		}

		ffare := FlightCachingFares{}
		if fareID != nil {
			tx.Where("id = ?", fareID).Take(&ffare)
		}

		opsigoClass := OpsigoFlightClassCaching{}
		tx.Where("flight_id = ?", data.TrxID).
			Where("class_id = ?", ffare.TrxID).
			Where("session_id = ?", data.SessionID).
			Take(&opsigoClass)

		if opsigoClass.FareBasisCode != nil {
			fbs := strings.Split(*opsigoClass.FareBasisCode, "|")
			opsigoClass.FareBasisCode = lib.Strptr(fbs[0])
			return opsigoClass.FareBasisCode
		}
	}

	return nil
}
