package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerCorporateGroup Integration Partner CorporateGroup
type IntegrationPartnerCorporateGroup struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerCorporateGroupAPI
	CorporateGroup     *CorporateGroup     `json:"corporate_group,omitempty" gorm:"foreignKey:CorporateGroupID;references:ID"`         // Corporate Group
	IntegrationPartner *IntegrationPartner `json:"integration_partner,omitempty" gorm:"foreignKey:IntegrationPartnerID;references:ID"` // Integration Partner
}

// IntegrationPartnerCorporateGroupAPI Integration Partner CorporateGroup API
type IntegrationPartnerCorporateGroupAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" gorm:"type:varchar(36);not null;"`               // Integration Partner ID
	CorporateGroupID     *uuid.UUID `json:"corporate_group_id,omitempty" gorm:"type:varchar(36);not null;"`                   // Croporate Group ID
	PartnerUsername      *string    `json:"partner_username,omitempty" gorm:"type:varchar(256);not null" example:"bayubuana"` // Partner Username
	PartnerPassword      *string    `json:"partner_password,omitempty" gorm:"type:varchar(256)" example:"root123"`            // Partner Password
	Organization         *string    `json:"organization,omitempty" gorm:"type:varchar(256)" example:"buana group"`            // Organization
	Domain               *string    `json:"domain,omitempty" gorm:"type:varchar(256)" example:"admin"`                        // Domain
	PCC                  *string    `json:"pcc,omitempty" gorm:"type:varchar(16)" example:"jkt"`                              // PCC
	Ipcc                 *string    `json:"ipcc,omitempty" gorm:"type:varchar(16)" example:"jkt"`                             // IPCC
	SystemID             *string    `json:"system_id,omitempty" gorm:"type:varchar(16)"`                                      // System ID
	AgencyID             *string    `json:"agency_id,omitempty" gorm:"type:varchar(16)"`                                      // Agency ID
	IATANumber           *string    `json:"iata_number,omitempty" gorm:"type:varchar(16)" example:"12345"`                    // IATA Number
	CustomerNumber       *string    `json:"customer_number,omitempty" gorm:"type:varchar(16)" example:"12345"`                // Customer Number
	PrinterAddress       *string    `json:"printer_address,omitempty" gorm:"type:varchar(16)" example:"16.175.22.128"`        // Printer Address
}

// IntegrationPartnerCorporateGroupRequest model
type IntegrationPartnerCorporateGroupRequest struct {
	// CorporateGroup
	// IntegrationPartnerCorporateGroup
	OfficeID        *uuid.UUID `json:"office_id,omitempty" swaggertype:"string" format:"uuid" validate:"required"`                  // The reference to the office. This indicates the grouping is based on an office
	PartnerUsername *string    `json:"partner_username,omitempty" gorm:"-" example:"xyzsurabaya" validate:"required,gte=0,lte=256"` // Partner Username
	PartnerPassword *string    `json:"partner_password,omitempty" gorm:"-" example:"password" validate:"required,gte=0,lte=256"`    // Partner Password
	Organization    *string    `json:"organization,omitempty" gorm:"-" example:"XYZ Surabaya" validate:"required,gte=0,lte=256"`    // Organization
	Domain          *string    `json:"domain,omitempty" gorm:"-" example:"www.xyzsurabaya.com" validate:"required,gte=0,lte=256"`   // Domain
	PCC             *string    `json:"pcc,omitempty" gorm:"-" example:"F9Q8" validate:"required,gte=0,lte=16"`                      // PCC
	IPCC            *string    `json:"ipcc,omitempty" gorm:"-" example:"D8UD" validate:"required,gte=0,lte=16"`                     // IPCC
	SystemID        *string    `json:"system_id,omitempty" gorm:"-" swaggertype:"string" validate:"required,gte=0,lte=16"`          // System ID
	AgencyID        *string    `json:"agency_id,omitempty" gorm:"-" swaggertype:"string" validate:"required,gte=0,lte=16"`          // Agency ID
	IATANumber      *string    `json:"iata_number,omitempty" gorm:"-" example:"12345" validate:"required,gte=0,lte=16"`             // IATA Number
	CustomerNumber  *string    `json:"customer_number,omitempty" gorm:"-" example:"12345" validate:"gte=0,lte=16"`                  // Customer Number
	PrinterAddress  *string    `json:"printer_address,omitempty" gorm:"-" example:"16.175.22.128" validate:"required,gte=0,lte=16"` // Printer Address
}

// IntegrationPartnerCredentialSelfServiceRequest model
type IntegrationPartnerCredentialSelfServiceRequest struct {
	PartnerUsername *string `json:"partner_username,omitempty" gorm:"-" example:"xyzsurabaya" validate:"required,gte=0,lte=256"` // Partner Username
	PartnerPassword *string `json:"partner_password,omitempty" gorm:"-" example:"password" validate:"required,gte=0,lte=256"`    // Partner Password
	Organization    *string `json:"organization,omitempty" gorm:"-" example:"XYZ Surabaya" validate:"required,gte=0,lte=256"`    // Organization
	Domain          *string `json:"domain,omitempty" gorm:"-" example:"www.xyzsurabaya.com" validate:"required,gte=0,lte=256"`   // Domain
	PCC             *string `json:"pcc,omitempty" gorm:"-" example:"F9Q8" validate:"required,gte=0,lte=16"`                      // PCC
	IPCC            *string `json:"ipcc,omitempty" gorm:"-" example:"D8UD" validate:"required,gte=0,lte=16"`                     // IPCC
	SystemID        *string `json:"system_id,omitempty" gorm:"-" swaggertype:"string" validate:"required,gte=0,lte=16"`          // System ID
	AgencyID        *string `json:"agency_id,omitempty" gorm:"-" swaggertype:"string" validate:"required,gte=0,lte=16"`          // Agency ID
	IATANumber      *string `json:"iata_number,omitempty" gorm:"-" example:"12345" validate:"required,gte=0,lte=16"`             // IATA Number
	CustomerNumber  *string `json:"customer_number,omitempty" gorm:"-" example:"12345" validate:"gte=0,lte=16"`                  // Customer Number
	PrinterAddress  *string `json:"printer_address,omitempty" gorm:"-" example:"16.175.22.128" validate:"required,gte=0,lte=16"` // Printer Address
}

// IntegrationPartnerCorporateGroupList model
type IntegrationPartnerCorporateGroupList struct {
	ID *uuid.UUID `json:"id,omitempty"`
	CorporateGroupAPI
	IntegrationPartnerCorporateGroupAPI
	AddressID *uuid.UUID `json:"address_id,omitempty"`
	Address   *Address   `json:"address,omitempty"`
	OfficeID  *uuid.UUID `json:"office_id,omitempty"`
	Office    *Office    `json:"office,omitempty"`
}
