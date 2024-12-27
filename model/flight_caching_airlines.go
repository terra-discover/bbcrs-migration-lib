package model

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// FlightCachingAirlines model
type FlightCachingAirlines struct {
	basic.Base
	basic.DataOwner
	SessionID           *uuid.UUID                           `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_airlines_session_id;not null"`
	FlightCachingID     *uuid.UUID                           `json:"flight_caching_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_airlines_flight_caching_id;not null" format:"uuid"` // FlightCachingID
	TrxID               *string                              `json:"trx_id,omitempty"`                                                                                                               //unique 3rd party flight airlines id
	GroupAirlineCode    *string                              `json:"group_airline_code,omitempty"`
	ClassRef            *string                              `json:"class_ref,omitempty" gorm:"type:varchar(36)"`
	ClassCode           *string                              `json:"class_code,omitempty" gorm:"type:varchar(36)"`
	ClassName           *string                              `json:"class_name,omitempty" gorm:"type:varchar(36)"`
	AirlineCode         *string                              `json:"airline_code,omitempty" gorm:"type:varchar(36)"`  // TODO : immediately will be removed
	AirlineName         *string                              `json:"airline_name,omitempty" gorm:"type:varchar(255)"` // TODO : immediately will be removed
	Detail              *basic.CustomDataTypes               `json:"detail,omitempty"`
	IsMultipleAirline   *bool                                `json:"is_multiple_airline,omitempty"`
	StopCount           *int                                 `json:"stop_count,omitempty"`
	Index               *int                                 `json:"index,omitempty"`
	Group               *string                              `json:"group,omitempty"`
	TotalDuration       *int64                               `json:"total_duration,omitempty"`
	TotalStringDuration *string                              `json:"total_string_duration,omitempty"`
	DepartureCityCode   *string                              `json:"departure_city_code,omitempty"`
	DepartureCityName   *string                              `json:"departure_city_name,omitempty"`
	ArrivalCityCode     *string                              `json:"arrival_city_code,omitempty"`
	ArrivalCityName     *string                              `json:"arrival_city_name,omitempty"`
	FlightCachingRoutes *[]FlightCachingRoutes               `json:"routes,omitempty" gorm:"foreignKey:FlightCachingAirlinesID;references:ID"`
	Violations          []FlightCachingTravelPolicyViolation `json:"violations"`
}

// GetFlightCachingAirlines by query filter
func (data *FlightCachingAirlines) GetFlightCachingAirlines(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
}

func (data *FlightCachingAirlines) GetGroupAirlineCode() (airlineCode int) {
	if data.GroupAirlineCode != nil {
		ga := *data.GroupAirlineCode
		airlineCode, _ = strconv.Atoi(ga)
	}

	if airlineCode == 0 {
		a, _ := json.Marshal(data)
		log.Printf("ERR OPSIGO BOOK getAirlineCode : Fail to get grouping airline code from airline instance : %s", a)
	}

	return
}
