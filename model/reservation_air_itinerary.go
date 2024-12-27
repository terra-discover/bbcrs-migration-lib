package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// ReservationAirItinerary Reservation Air Itinerary
type ReservationAirItinerary struct {
	basic.Base
	basic.DataOwner
	ReservationAirItineraryAPI
	Reservation  *Reservation  `json:"reservation,omitempty" gorm:"foreignKey:ReservationID"`
	AirItinerary *AirItinerary `json:"air_itinerary,omitempty" gorm:"foreignKey:AirItineraryID"`
}

// ReservationAirItineraryAPI Reservation Air Itinerary API
type ReservationAirItineraryAPI struct {
	ReservationID  *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"`   // Reservation ID
	AirItineraryID *uuid.UUID `json:"air_itinerary_id,omitempty" gorm:"type:varchar(36);uniqueIndex:,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Air Itinerary ID
}

// GetReservationAirItinerary by query filter
func (data *ReservationAirItinerary) GetReservationAirItinerary(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&data)
}
