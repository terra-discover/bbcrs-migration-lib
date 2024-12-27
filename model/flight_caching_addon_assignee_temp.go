package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Digunakan untuk fitur modify addon dan terkait IsModifyBookingRequest = true
type FlightCachingAddonAssigneeTemp struct {
	basic.Base
	basic.DataOwner
	SessionID            *uuid.UUID         `json:"session_id,omitempty"`
	FlightCachingAddonID *uuid.UUID         `json:"flight_caching_addon_id,omitempty"`
	FlightTravellerID    *uuid.UUID         `json:"flight_traveller_id,omitempty"`
	PersonID             *uuid.UUID         `json:"person_id,omitempty"`
	FlightCachingAddon   FlightCachingAddon `json:"flight_caching_addon"`
	ProductID            *uuid.UUID         `json:"product_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`
	ProductIDReservation *uuid.UUID         `json:"product_id_reservation,omitempty" gorm:"type:varchar(36)" format:"uuid"` // productID reservation
	IsPaid               *bool              `json:"is_paid,omitempty"`                                                      // IsPaid
	IsApproved           *bool              `json:"is_approved,omitempty"`                                                  // untuk trigger apakah temporary ini udah ada trx (create, delete) atau belum , false belum ada trx (create, delete)
}
