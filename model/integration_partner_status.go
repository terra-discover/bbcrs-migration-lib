package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// IntegrationPartnerStatus struct
type IntegrationPartnerStatus struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerStatusAPI
	//IntegrationPartner *IntegrationPartner
}

// IntegrationPartnerStatusAPI struct
type IntegrationPartnerStatusAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null"`                              // IntegrationPartnerID
	StatusID             *uuid.UUID `json:"status_id,omitempty" gorm:"type:varchar(36);not null"`                                           // StatusID
	StatusCode           *string    `json:"status_code,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null"` // StatusCode
	StatusName           *string    `json:"status_name,omitempty" gorm:"type:varchar(512);index:,unique,where:deleted_at is null"`          // StatusName
}

// IntegrationPartnerStatuses struct
type IntegrationPartnerStatuses []IntegrationPartnerStatus

// Seed data
func (integrationPartnerStatuses *IntegrationPartnerStatuses) Seed() *IntegrationPartnerStatuses {
	code := []string{"OPEN", "USED", "VOID", "PRTD", "EXCH", "RFND", "CKIN", "CTRL", "ACTL", "SUSP", "OK", "REAC"}
	name := []string{"Ticket open, unused", "Ticket used, lifted/boarded", "Ticket void, transaction voided", "Ticket printed, flight coupons printed", "Ticket exchanged, exchanged/reissued",
		"Ticket refunded", "Ticket checked-in", "Ticket under airport control", "Ticket under airport control", "Ticket suspended by carrier", "Ticket OK, okay for travel", "Ticket reactivated, open, when void an exchange"}
	for i := 0; i < len(code); i++ {
		*integrationPartnerStatuses = append(*(integrationPartnerStatuses), IntegrationPartnerStatus{
			IntegrationPartnerStatusAPI: IntegrationPartnerStatusAPI{
				StatusCode:           lib.Strptr(code[i]),
				StatusName:           lib.Strptr(name[i]),
				IntegrationPartnerID: lib.UUIDPtr(uuid.New()),
				StatusID:             lib.UUIDPtr(uuid.New()),
			},
		})
	}
	return integrationPartnerStatuses
}
