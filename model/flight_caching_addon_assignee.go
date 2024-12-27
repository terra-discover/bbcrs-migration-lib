package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

type FlightCachingAddonAssignee struct {
	basic.Base
	basic.DataOwner
	SessionID              *uuid.UUID              `json:"session_id,omitempty"`
	FlightCachingAddonID   *uuid.UUID              `json:"flight_caching_addon_id,omitempty"`
	FlightTravellerID      *uuid.UUID              `json:"flight_traveller_id,omitempty"`
	PersonID               *uuid.UUID              `json:"person_id,omitempty"`
	FlightCachingAddon     *FlightCachingAddon     `json:"flight_caching_addon"`
	ProductID              *uuid.UUID              `json:"product_id,omitempty" gorm:"type:varchar(36)" format:"uuid"`             // productID proposal
	ProductIDReservation   *uuid.UUID              `json:"product_id_reservation,omitempty" gorm:"type:varchar(36)" format:"uuid"` // productID reservation
	IsCancelled            *bool                   `json:"is_cancelled,omitempty"`                                                 // IsCancelled
	IsPaid                 *bool                   `json:"is_paid,omitempty"`                                                      // IsPaid
	FlightCachingTraveller *FlightCachingTraveller `json:"traveller,omitempty" gorm:"foreignKey:FlightTravellerID"`
}

// BeforeCreate Data
func (b *FlightCachingAddonAssignee) BeforeCreate(tx *gorm.DB) error {
	b.Base.BeforeCreate(tx)
	if nil == b.IsCancelled {
		b.IsCancelled = lib.Boolptr(false)
	}
	return nil
}
