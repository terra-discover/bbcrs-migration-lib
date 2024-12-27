package model

import (
	"log"
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FlightCachingRequest model
type FlightCachingRequest struct {
	basic.Base
	basic.DataOwner
	basic.AuthContext
	SessionID              *uuid.UUID                        `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_request_session_id;not null"`
	ExpiredIn              *int64                            `json:"expired_in,omitempty"`
	Adult                  *int                              `json:"adult,omitempty"`
	CabinTypeCode          *string                           `json:"cabin_type_code,omitempty" example:"ECO"`         // [DEPRECATED] get lists from "cabin_type.cabin_type_code"
	CabinTypeID            *uuid.UUID                        `json:"cabin_type_id,omitempty" gorm:"type:varchar(36)"` //will be filled by logic based on CabinTypeCode
	Carriers               *string                           `json:"carriers,omitempty"`                              // multiple value (airline_code) separates by comma
	Child                  *int                              `json:"child,omitempty"`
	DepartingArrivalFrom   *string                           `json:"departing_arrival_from,omitempty" gorm:"type:varchar(50)"`
	DepartingArrivalTo     *string                           `json:"departing_arrival_to,omitempty" gorm:"type:varchar(50)"`
	DepartingDepartureFrom *string                           `json:"departing_departure_from,omitempty" gorm:"type:varchar(50)"`
	DepartingDepartureTo   *string                           `json:"departing_departure_to,omitempty" gorm:"type:varchar(50)"`
	Infant                 *int                              `json:"infant,omitempty"`
	IsPersonalTrip         *bool                             `json:"is_personal_trip,omitempty"`
	IsSelfBookingTool      *bool                             `json:"is_self_booking_tool,omitempty"`
	ReturningArrivalFrom   *string                           `json:"returning_arrival_from,omitempty" gorm:"type:varchar(50)"`
	ReturningArrivalTo     *string                           `json:"returning_arrival_to,omitempty" gorm:"type:varchar(50)"`
	ReturningDepartureFrom *string                           `json:"returning_departure_from,omitempty" gorm:"type:varchar(50)"`
	ReturningDepartureTo   *string                           `json:"returning_departure_to,omitempty" gorm:"type:varchar(50)"`
	TripTypeCode           *string                           `json:"trip_type_code,omitempty" gorm:"type:varchar(50)"`
	IsDomestic             *bool                             `json:"is_domestic,omitempty"`
	Flights                *[]FlightCachingRequestRoute      `json:"flights,omitempty" gorm:"foreignKey:FlightRequestID;references:ID"`
	Travellers             *[]FlightCachingRequestTravellers `json:"travellers,omitempty" gorm:"foreignKey:FlightRequestID;references:ID"`
}

// GetTripType by query filter
func (data *FlightCachingRequest) GetFlightCachingRequest(tx *gorm.DB, queryFilters string) *FlightCachingRequest {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Preload(`Flights`).Preload(`Travellers`).Take(&data)
	return data
}

func (data *FlightCachingRequest) GetDestinationGroupAffected(db *gorm.DB) (destinationGroup []DestinationGroup) {
	_, country := data.GetCountryCityAffected(db)

	countryIDs := []string{}
	for _, ctr := range country {
		countryIDs = append(countryIDs, ctr.ID.String())
	}

	db.Model(&DestinationGroup{}).
		Select(`destination_group.*`).
		Where(`EXISTS (
		SELECT dgc.id FROM destination_group_country dgc WHERE dgc.destination_group_id = destination_group.id AND dgc.country_id IN ?
		)`, countryIDs).
		Find(&destinationGroup)

	return
}

func (data *FlightCachingRequest) GetDestinationGroupMapCountry(db *gorm.DB) (mapCountryDestinationGroup map[string][]DestinationGroup) {
	_, country := data.GetCountryCityAffected(db)

	mapCountryDestinationGroup = make(map[string][]DestinationGroup)

	countryIDs := []string{}
	for _, ctr := range country {
		countryIDs = append(countryIDs, ctr.ID.String())
	}

	destinationGroup := []DestinationGroup{}
	db.Model(&DestinationGroup{}).
		Preload("DestinationGroupCountry").
		Select(`destination_group.*`).
		Where(`EXISTS (
		SELECT dgc.id FROM destination_group_country dgc WHERE dgc.destination_group_id = destination_group.id AND dgc.country_id IN ?
		)`, countryIDs).
		Find(&destinationGroup)

	cids := []string{}
	for _, dg := range destinationGroup {
		for _, dgc := range dg.DestinationGroupCountry {
			if dgc.CountryID != nil {
				cid := dgc.CountryID.String()
				if lib.SliceContains(countryIDs, cid) {
					mapCountryDestinationGroup[cid] = append(mapCountryDestinationGroup[cid], dg)
					cids = append(cids, cid)
				}
			}
		}
	}

	log.Printf("%+v", cids)

	return
}

func (data *FlightCachingRequest) GetCountryCityAffected(db *gorm.DB) (cities []City, countries []Country) {
	cityCode := []string{}

	if data.Flights == nil {
		return
	}

	for _, route := range *data.Flights {
		// read destination group : only read the arrival from request route
		if route.ArrivalCityCode != nil {
			if exists := lib.SliceContains(cityCode, *route.ArrivalCityCode); !exists {
				cityCode = append(cityCode, *route.ArrivalCityCode)
			}
		}
	}

	var countryIDHelper []string
	if len(cityCode) > 0 {
		db.Preload("Country").Where("city_code IN ?", cityCode).Find(&cities)
		for _, cityData := range cities {
			if cityData.Country != nil {
				if exists := lib.SliceContains(countryIDHelper, cityData.Country.ID.String()); !exists {
					countryIDHelper = append(countryIDHelper, cityData.Country.ID.String())
					countries = append(countries, *cityData.Country)
				}
			}
		}
	}

	return
}

func (f *FlightCachingRequest) IsDomesticRoutes(db *gorm.DB) (isDomestic bool) {
	if f.Flights == nil {
		return false
	}
	if nil != f && len(*f.Flights) > 0 {
		for i := range *f.Flights {
			item := (*f.Flights)[i]
			dCity := ""
			aCity := ""
			if nil != item.DepartureCityCode {
				dCity = strings.ToLower(*item.DepartureCityCode)
			}
			if nil != item.ArrivalCityCode {
				aCity = strings.ToLower(*item.ArrivalCityCode)
			}

			var total int64
			db.Model(&Country{}).
				Joins(`INNER JOIN "city" "dc" ON "dc".country_id = "country".id`).
				// Joins(`INNER JOIN "airport" "da" ON "da".city_id = "dc".id`).
				Joins(`INNER JOIN "city" "ac" ON "ac".country_id = "country".id`).
				// Joins(`INNER JOIN "airport" "aa" ON "aa".city_id = "ac".id`).
				Where(`LOWER("dc".city_code) = ? `, dCity).
				Where(`LOWER("ac".city_code) = ? `, aCity).
				Where(`LOWER("country".country_name) = ?`, "indonesia").
				Count(&total)

			if total == 0 {
				return false
			}
		}

		return true
	}

	return false
}
