package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// AirTraveler Air Traveler
type AirTraveler struct {
	basic.Base
	basic.DataOwner
	AirTravelerAPI
	Employee *Employee `json:"employee,omitempty" gorm:"foreignKey:EmployeeID"`
	Person   *Person   `json:"person,omitempty" gorm:"foreignKey:PersonID"`
}

// AirTravelerAPI Air Traveler API
type AirTravelerAPI struct {
	AirItineraryID          *uuid.UUID  `json:"air_itinerary_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Air Itinerary ID
	TravelerIndexNumber     *int        `json:"traveler_index_number,omitempty" gorm:"type:smallint" swaggertype:"integer"`                     // Traveler Index Number
	IsPrimary               *bool       `json:"is_primary,omitempty" gorm:"type:boolean" swaggertype:"boolean"`                                 // Is Primary
	PassengerTypeID         *uuid.UUID  `json:"passenger_type_id" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`          // Passenger Type ID
	Age                     *int        `json:"age,omitempty" gorm:"type:smallint" swaggertype:"integer"`                                       // Age
	PersonID                *uuid.UUID  `json:"person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                 // Person ID
	EmployeeID              *uuid.UUID  `json:"employee_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`               // Employee ID
	AccompaniedByInfant     *bool       `json:"accompanied_by_infant,omitempty" gorm:"type:boolean" swaggertype:"boolean"`                      // Accompanied By Infant
	ETicketIndicator        *bool       `json:"e_ticket_indicator,omitempty" gorm:"type:boolean" swaggertype:"string"`                          // E Ticket Indicator
	FrequentFlyerNumber     *string     `json:"frequent_flyer_number,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`                   // Frequent Flyer Number
	SpecialRequest          *string     `json:"special_request,omitempty" gorm:"type:text" swaggertype:"string"`                                // Special Request
	Comment                 *string     `json:"comment,omitempty" gorm:"type:text" swaggertype:"string"`                                        // Comment
	ConsumerViewableComment *string     `json:"consumer_viewable_comment,omitempty" gorm:"type:text" swaggertype:"string"`
	Additional              interface{} `json:"additional,omitempty" gorm:"-"`
}

// GetAirTraveler by query filter
func (data *AirTraveler) GetAirTraveler(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Take(&data)
}
