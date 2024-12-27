package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ReservationIdentifier Reservation Identifier
type ReservationIdentifier struct {
	basic.Base
	basic.DataOwner
	UniqueIDType *UniqueIDType `json:"unique_id_type,omitempty"`
	ReservationIdentifierAPI
}

// ReservationIdentifierAPI Reservation Identifier API
type ReservationIdentifierAPI struct {
	ReservationID         *uuid.UUID `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`    // Description: The reference to the reservation.
	UniqueIDTypeID        *uuid.UUID `json:"unique_id_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // The value of the unique identifier.
	UniqueIDValue         *string    `json:"unique_id_value,omitempty" gorm:"type:varchar(64);not null"`                                      // The value of the unique identifier.
	BusinessPartnerID     *uuid.UUID `json:"business_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`        // The reference to the business partner.
	DestinationSystemID   *uuid.UUID `json:"destination_system_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`      // The reference to the destination system
	DistributionPartnerID *uuid.UUID `json:"distribution_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`    // The reference to the distribution partner
	IntegrationPartnerID  *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`     // The reference to the integration partner.
	SupplierID            *uuid.UUID `json:"supplier_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                // The reference to the supplier, e.g. hotel supplier, airline
	ReferenceID           *uuid.UUID `json:"reference_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`               // A reference to an air itinerary, or air origin destination option, or flight segment, or room stay.
	IsVerified            *bool      `json:"is_verified,omitempty"`                                                                           // Indicates whether the air ticket has been verified, e.g. has been found in the daily sales report.
	IsPrimary             *bool      `json:"is_primary,omitempty"`                                                                            // Is Primary
}
