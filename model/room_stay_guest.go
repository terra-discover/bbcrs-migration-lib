package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RoomStayGuest model
type RoomStayGuest struct {
	basic.Base
	basic.DataOwner
	RoomStayGuestAPI
	Person   *Person   `json:"person,omitempty"`
	Employee *Employee `json:"employee,omitempty"`
}

// RoomStayGuestAPI struct
type RoomStayGuestAPI struct {
	RoomStayID           *uuid.UUID `json:"room_stay_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                    // Room Stay ID
	RoomIndexNumber      *int       `json:"room_index_number,omitempty" gorm:"type:smallint"`                                                     // Room Index Number
	GuestIndexNumber     *int       `json:"guest_index_number,omitempty" gorm:"type:smallint"`                                                    // Guset Index Number
	IsPrimary            *bool      `json:"is_primary,omitempty" gorm:"type:bool"`                                                                // Is Primary
	AgeQualifyingTypeID  *uuid.UUID `json:"age_qualifying_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Age Qualifying Type ID
	Age                  *int       `json:"age,omitempty" gorm:"type:smallint"`                                                                   // Age
	PersonID             *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                       // Person ID
	EmployeeID           *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                     // Employee ID
	ArrivalTime          *string    `json:"arrival_time,omitempty" gorm:"type:time(0)"`                                                           // Arrival Time
	DepartureTime        *string    `json:"departure_time,omitempty" gorm:"type:time(0)"`                                                         // Departure Time
	ArrivalTransportID   *uuid.UUID `json:"arrival_transport_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`            // Arrival Transport ID
	DepartureTransportID *uuid.UUID `json:"departure_transport_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`          // Departure Transport ID
	SpecialRequest       *string    `json:"special_request,omitempty" gorm:"type:text"`                                                           // Special Request
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                   // Comment
	GuestViewableComment *string    `json:"guest_viewable_comment,omitempty" gorm:"type:text"`                                                    // Guest Viewable Comment
}
