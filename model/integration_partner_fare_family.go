package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// IntegrationPartnerFareFamily Integration Partner FareFamily
type IntegrationPartnerFareFamily struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerFareFamilyAPI
	CabinType *CabinType `json:"cabin_type,omitempty"`
	FareType  *FareType  `json:"fare_type,omitempty"`
}

// IntegrationPartnerFareFamilyAPI Integration PartnerFareFamily API
// In one integration_partner_id only can contains unique booking_class
type IntegrationPartnerFareFamilyAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`
	CabinTypeID          *uuid.UUID `json:"cabin_type_id,omitempty" gorm:"type:varchar(36);not null;"`
	FareTypeID           *uuid.UUID `json:"fare_type_id,omitempty" gorm:"type:varchar(36);not null;"`
	BookingClass         *string    `json:"booking_class,omitempty" gorm:"type:varchar(2);not null;"`
}

// GetIntegrationPartnerFareFamily
func (s *IntegrationPartnerFareFamily) Get(tx *gorm.DB, queryFilters string) *[]IntegrationPartnerFareFamily {
	result := []IntegrationPartnerFareFamily{}
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Model(&result).Where(qf, wf...).Preload(`CabinType`).Preload(`FareType`).Find(&result)
	return &result
}
