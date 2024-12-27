package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerHotelSupplier Integration Partner HotelSupplier
type IntegrationPartnerHotelSupplier struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerHotelSupplierAPI
	HotelSupplier      *HotelSupplier      `json:"hotel_supplier,omitempty"`
	IntegrationPartner *IntegrationPartner `json:"integration_partner,omitempty"`
}

// IntegrationPartnerHotelSupplierAPI Integration PartnerHotelSupplier API
type IntegrationPartnerHotelSupplierAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`
	HotelSupplierID      *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);"`
	HotelSupplierCode    *string    `json:"hotel_supplier_code,omitempty" gorm:"type:varchar(36);not null"`
	HotelSupplierName    *string    `json:"hotel_supplier_name,omitempty" gorm:"type:varchar(256);"`
}
