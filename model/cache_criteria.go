package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CacheCriteria Cache Criteria
type CacheCriteria struct {
	basic.Base
	basic.DataOwner
	AgentID *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36)" format:"uuid" validate:"required"` // agent id
	CacheCriteriaAPI
	CacheAirOriginDestinationCriteria []*CacheAirOriginDestinationCriteria `json:"cache_air_origin_destination_criteria,omitempty" gorm:"foreignKey:CacheCriteriaID"`                          // cache air origin destination criteria
	CacheAirTravelPreferenceCriteria  *CacheAirTravelPreferenceCriteria    `json:"cache_air_travel_preference_criteria,omitempty" gorm:"foreignKey:ID;references:CacheCriteriaID"`             // cache air travel preference criteria
	CacheAirTravelerCriteria          *CacheAirTravelerCriteria            `json:"cache_air_traveler_criteria,omitempty" gorm:"foreignKey:ID;references:CacheCriteriaID" swaggerignore:"true"` // cache air traveler criteria
	CacheGuestCriteria                *CacheGuestCriteria                  `json:"cache_guest_criteria,omitempty" gorm:"foreignKey:ID;references:CacheCriteriaID"`                             // cache guest criteria
	CacheRoomStayCriteria             *CacheRoomStayCriteria               `json:"cache_room_stay_criteria,omitempty" gorm:"foreignKey:ID;references:CacheCriteriaID"`                         // cache room stay criteria
	Agent                             *Agent                               `json:"agent,omitempty" gorm:"foreignKey:AgentID;references:ID"`                                                    // agent
	Corporate                         *Corporate                           `json:"corporate,omitempty" gorm:"foreignKey:CorporateID;references:ID"`                                            // corporate
	TripType                          *TripType                            `json:"trip_type,omitempty" gorm:"foreignKey:TripTypeID;references:ID"`                                             // trip type
	Currency                          *Currency                            `json:"currency,omitempty" gorm:"foreignKey:CurrencyID;references:ID"`                                              // currency
	Language                          *Language                            `json:"language,omitempty" gorm:"foreignKey:LanguageID;references:ID"`                                              // language
}

// CacheCriteriaAPI Cache Criteria API
type CacheCriteriaAPI struct {
	CorporateID                       *uuid.UUID                             `json:"corporate_id,omitempty" gorm:"type:varchar(36)" format:"uuid" validate:"omitempty,required"` // corporate id
	TripTypeID                        *uuid.UUID                             `json:"trip_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid" validate:"omitempty,required"` // trip type id
	CurrencyID                        *uuid.UUID                             `json:"currency_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                // currency id
	LanguageID                        *uuid.UUID                             `json:"language_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`                                // language id
	CacheAirOriginDestinationCriteria []CacheAirOriginDestinationCriteriaAPI `json:"cache_air_origin_destination_criteria,omitempty" gorm:"-"`
	CacheAirTravelPreferenceCriteria  CacheAirTravelPreferenceCriteriaAPI    `json:"cache_air_travel_preference_criteria,omitempty" gorm:"-"`
	CacheAirTravelerCriteria          CacheAirTravelerCriteriaAPI            `json:"cache_air_traveler_criteria,omitempty" gorm:"-"`
	CacheGuestCriteria                CacheGuestCriteriaAPI                  `json:"cache_guest_criteria,omitempty" gorm:"-"`
	CacheRoomStayCriteria             CacheRoomStayCriteriaAPI               `json:"cache_room_stay_criteria,omitempty" gorm:"-"`
}

// CacheCriteriaFlightGetData Cache Criteria Flight Get Data
type CacheCriteriaFlightGetData struct {
	CacheCriteria
	OriginLocations      *string    `json:"origin_locations,omitempty"`
	DestinationLocations *string    `json:"destination_locations,omitempty"`
	DepartureDatetimes   *string    `json:"departure_datetimes,omitempty"`
	ArrivalDatetimes     *string    `json:"arrival_datetimes,omitempty"`
	OriginAirports       *string    `json:"origin_airports,omitempty"`
	DestinationAirports  *string    `json:"destination_airports,omitempty"`
	NumberOfTravelers    *int       `json:"number_of_travelers,omitempty"`
	CabinType            *CabinType `json:"cabin_type,omitempty"`
	CabinTypeID          *uuid.UUID `json:"cabin_type_id,omitempty"`
}

// CacheCriteriaHotelGetData Cache Criteria Hotel Get Data
type CacheCriteriaHotelGetData struct {
	CacheCriteria
	DestinationLocations *string `json:"destination_locations,omitempty"`
	CheckIns             *string `json:"checkins,omitempty"`
	DestinationHotels    *string `json:"destination_hotels,omitempty"`
	DestinationCities    *string `json:"destination_cities,omitempty"`
}

// CacheCriteriaPostHotelData Cache Criteria API
type CacheCriteriaPostHotelData struct {
	AgentID               *uuid.UUID               `json:"agent_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`     // agent id
	CorporateID           *uuid.UUID               `json:"corporate_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // corporate id
	TripTypeID            *uuid.UUID               `json:"trip_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // trip type id
	CurrencyID            *uuid.UUID               `json:"currency_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`  // currency id
	LanguageID            *uuid.UUID               `json:"language_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`  // language id
	CacheGuestCriteria    CacheGuestCriteriaAPI    `json:"cache_guest_criteria,omitempty" gorm:"-"`
	CacheRoomStayCriteria CacheRoomStayCriteriaAPI `json:"cache_room_stay_criteria,omitempty" gorm:"-"`
}

// CacheCriteriaPostFlightData Cache Criteria API
type CacheCriteriaPostFlightData struct {
	AgentID                           *uuid.UUID                             `json:"agent_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`     // agent id
	CorporateID                       *uuid.UUID                             `json:"corporate_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // corporate id
	TripTypeID                        *uuid.UUID                             `json:"trip_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // trip type id
	CurrencyID                        *uuid.UUID                             `json:"currency_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`  // currency id
	LanguageID                        *uuid.UUID                             `json:"language_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`  // language id
	CacheAirOriginDestinationCriteria []CacheAirOriginDestinationCriteriaAPI `json:"cache_air_origin_destination_criteria,omitempty" gorm:"-"`
	CacheAirTravelPreferenceCriteria  CacheAirTravelPreferenceCriteriaAPI    `json:"cache_air_travel_preference_criteria,omitempty" gorm:"-"`
	CacheAirTravelerCriteria          CacheAirTravelerCriteriaAPI            `json:"cache_air_traveler_criteria,omitempty" gorm:"-"`
}
