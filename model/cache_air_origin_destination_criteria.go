package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CacheAirOriginDestinationCriteria Cache Air Origin Destination Criteria
type CacheAirOriginDestinationCriteria struct {
	basic.Base
	basic.DataOwner
	CacheAirOriginDestinationCriteriaAPI
	CacheCriteria      *CacheCriteria `json:"cache_criteria,omitempty" gorm:"foreignKey:CacheCriteriaID;references:ID" swaggerignore:"true"` // cache criteria
	OriginCity         *City          `json:"origin_city,omitempty" gorm:"foreignKey:OriginCityID;references:ID"`                            // origin city
	DestinationCity    *City          `json:"destination_city,omitempty" gorm:"foreignKey:DestinationCityID;references:ID"`                  // destination city
	OriginAirport      *Airport       `json:"origin_airport,omitempty" gorm:"foreignKey:OriginAirportID;references:ID"`                      // origin airport
	DestinationAirport *Airport       `json:"destination_airport,omitempty" gorm:"foreignKey:DestinationAirportID;references:ID"`            // destination airport
}

// CacheAirOriginDestinationCriteriaAPI  Cache Air Origin Destination Criteria API
type CacheAirOriginDestinationCriteriaAPI struct {
	CacheCriteriaID          *uuid.UUID       `json:"cache_criteria_id,omitempty" gorm:"type:varchar(36);not null" format:"uuid" swaggerignore:"true"`   // cache criteria id
	IndexNumber              *int             `json:"index_number,omitempty" gorm:"type:smallint;" example:"1"`                                          // index number
	OriginLocation           *string          `json:"origin_location,omitempty" gorm:"type:varchar(36)"`                                                 // origin location
	DestinationLocation      *string          `json:"destination_location,omitempty" gorm:"type:varchar(36)"`                                            // destination location
	OriginCityID             *uuid.UUID       `json:"origin_city_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                    // origin city id
	OriginAirportID          *uuid.UUID       `json:"origin_airport_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                 // origin airport id
	DestinationCityID        *uuid.UUID       `json:"destination_city_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                               // destination city id
	DestinationAirportID     *uuid.UUID       `json:"destination_airport_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                            // destination airport id
	DepartureDateTime        *strfmt.DateTime `json:"departure_datetime,omitempty" gorm:"type:timestamptz;null" swaggertype:"string" format:"date-time"` // departure datetime
	ArrivalDateTime          *strfmt.DateTime `json:"arrival_datetime,omitempty" gorm:"type:timestamptz;null" swaggertype:"string" format:"date-time"`   // arrival datetime
	CityID                   *uuid.UUID       `json:"city_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
	AirportID                *uuid.UUID       `json:"airport_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
	AdministrativeCityID     *uuid.UUID       `json:"administrative_city_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
	ZoneID                   *uuid.UUID       `json:"zone_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
	DestCityID               *uuid.UUID       `json:"dest_city_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
	DestAirportID            *uuid.UUID       `json:"dest_airport_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
	DestAdministrativeCityID *uuid.UUID       `json:"dest_administrative_city_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
	DestZoneID               *uuid.UUID       `json:"dest_zone_id,omitempty" gorm:"-" format:"uuid" swaggerignore:"true"`
}
