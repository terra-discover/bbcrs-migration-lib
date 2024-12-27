package model

import (
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// ReservationProfile Reservation Profile
type ReservationProfile struct {
	basic.Base
	basic.DataOwner
	ReservationProfileAPI
	ReservationAsset *ReservationAsset `json:"reservation_asset,omitempty" gorm:"foreignKey:ReservationID;references:ReservationID"` // Reservation Asset
	ProfileType      *ProfileType      `json:"profile_type,omitempty" gorm:"foreignKey:ProfileTypeID;references:ID"`
	Person           *Person           `json:"person,omitempty"`
	Employee         *Employee         `json:"employee,omitempty"`
}

// ReservationProfileAPI Reservation Profile API
type ReservationProfileAPI struct {
	ReservationID           *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`  // The reference to the reservation.
	ProfileTypeID           *uuid.UUID `json:"profile_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // The reference to the profile type, e.g. travel consultant, booker, traveler, etc.
	PersonID                *uuid.UUID `json:"person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                // The reference to the person.
	EmployeeID              *uuid.UUID `json:"employee_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // The reference to the employee.
	UserAccountID           *uuid.UUID `json:"user_account_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid "`         // The reference to the user.
	DivisionID              *uuid.UUID `json:"division_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // The reference to the division.
	OfficeID                *uuid.UUID `json:"office_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                // The reference to the office.
	TravelerIndexNumber     *int       `json:"traveler_index_number,omitempty" gorm:"type:smallint"`                                          // The reference to this traveler within the air itinerary.
	IsPrimary               *bool      `json:"is_primary,omitempty"`                                                                          // Indicates whether this is the primary traveler.
	PassengerTypeID         *uuid.UUID `json:"passenger_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`        // Defines an age range or age category of a passenger (e.g. adult, child, infant).
	AgeQualifyingTypeID     *uuid.UUID `json:"age_qualifying_type_id,omitempty"  gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`  // Defines an age range or age category of a guest (e.g. adult, child).
	Age                     *int       `json:"age,omitempty" gorm:"type:smallint"`                                                            // The reference to this traveler within the air itinerary.
	OnLapPersonID           *uuid.UUID `json:"on_lap_person_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`         // The reference to the person the infant is on lap.
	ApprovedOverCredit      *bool      `json:"approved_over_credit,omitempty"`                                                                // Indicates whether this person has approved over credit.
	SpecialRequest          *string    `json:"special_request,omitempty" gorm:"type:text"`                                                    // The special request of the proposal.
	Comment                 *string    `json:"comment,omitempty" gorm:"type:text"`                                                            // The internal comment of the proposal.
	ConsumerViewableComment *string    `json:"consumer_viewable_comment,omitempty" gorm:"type:text"`                                          //  The comment that is shown to the consumer.
}
