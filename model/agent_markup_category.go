package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentMarkupCategory Agent Markup Category
type AgentMarkupCategory struct {
	basic.Base
	basic.DataOwner
	AgentMarkupCategoryAPI
	MarkupCategory *MarkupCategory `json:"markup_category" gorm:"foreignKey:MarkupCategoryID;references:ID"`
	Agent          *Agent          `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentMarkupCategoryAPI Agent Markup Category Api
type AgentMarkupCategoryAPI struct {
	AgentID          *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	MarkupCategoryID *uuid.UUID `json:"markup_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}

// AgentMarkupCategoryDetail struct
type AgentMarkupCategoryDetail struct {
	basic.Base
	basic.DataOwner
	AgentMarkupCategoryDetailAPI
	Used                       *bool       `json:"used,omitempty"`
	NumberOfOverride           *int64      `json:"number_of_override,omitempty"`
	DomesticFlightMarkup       *MarkupRate `json:"domestic_flight_markup,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalFlightMarkup  *MarkupRate `json:"international_flight_markup,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticHotelMarkup        *MarkupRate `json:"domestic_hotel_markup,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalHotelMarkup   *MarkupRate `json:"international_hotel_markup,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticFlightMarkups      *string     `json:"domestic_flight_markups,omitempty"`
	InternationalFlightMarkups *string     `json:"international_flight_markups,omitempty"`
	DomesticHotelMarkups       *string     `json:"domestic_hotel_markups,omitempty"`
	InternationalHotelMarkups  *string     `json:"international_hotel_markups,omitempty"`
}

// AgentMarkupCategoryDetailAPI Agent Markup Category Detail Api
type AgentMarkupCategoryDetailAPI struct {
	MarkupCategoryName        *string             `json:"markup_category_name" example:"Flight Mark-up 1" validate:"required"`
	Description               *string             `json:"description" example:"Flight Mark-up 1"`
	DomesticFlightMarkup      *AgentMarkupRateAPI `json:"domestic_flight_markup" gorm:"-"`
	InternationalFlightMarkup *AgentMarkupRateAPI `json:"international_flight_markup" gorm:"-"`
	DomesticHotelMarkup       *AgentMarkupRateAPI `json:"domestic_hotel_markup" gorm:"-"`
	InternationalHotelMarkup  *AgentMarkupRateAPI `json:"international_hotel_markup" gorm:"-"`
}

// AgentMarkupRateAPI writeable struct
type AgentMarkupRateAPI struct {
	ChargeTypeID   *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // Charge Type Id
	Percent        *float64   `json:"percent,omitempty" gorm:"type:decimal(19,4)"`                    // Percent
	Amount         *float64   `json:"amount,omitempty" gorm:"type:decimal(19,4)"`                     // Amount
	IsTaxInclusive *bool      `json:"is_tax_inclusive,omitempty"`                                     // Is Tax Inclusive
}

// AgentMarkupCategoryMarkupRateDetail Agent Markup Category Markup Rate Detail
type AgentMarkupCategoryMarkupRateDetail struct {
	basic.Base
	basic.DataOwner
	Markup                *MarkupRate       `json:"markup,omitempty" gorm:"foreignKey:ID;references:ID"`
	DestinationGroup      *DestinationGroup `json:"destination_group,omitempty" gorm:"foreignKey:ID"`
	AirlineCategory       *AirlineCategory  `json:"airline_category,omitempty" gorm:"foreignKey:ID"`
	Airline               *Airline          `json:"airline,omitempty" gorm:"foreignKey:ID"`
	SupplierType          *SupplierType     `json:"supplier_type,omitempty" gorm:"foreignKey:ID"`
	HotelSupplier         *HotelSupplier    `json:"hotel_supplier,omitempty" gorm:"foreignKey:ID"`
	DestinationGroupNames *string           `json:"destination_group_names,omitempty"` // Destination Group Names joined by ampersand `,`
	AirlineNames          *string           `json:"airline_names,omitempty"`           // Airline Names joined by ampersand `,`
	AirlineCategoryNames  *string           `json:"airline_category_names,omitempty"`  // Airline Category Names joined by ampersand `,`
	SupplierTypeNames     *string           `json:"supplier_type_names,omitempty"`     // Supplier Type Names joined by ampersand `,`
	HotelSupplierNames    *string           `json:"hotel_supplier_names,omitempty"`    // Hotel Supplier Names joined by ampersand `,`
	Markups               *string           `json:"markups,omitempty"`                 // Markups joined by ampersand `,`
}

// AgentMarkupCategoryMarkupRateResultDetail Agent Markup Category Markup Rate Result Detail
type AgentMarkupCategoryMarkupRateResultDetail struct {
	basic.Base
	basic.DataOwner
	AgentMarkupCategoryMarkupRateDetailAPI
	Markup *MarkupRate `json:"markup,omitempty" gorm:"foreignKey:ID;references:ID"`
}

// AgentMarkupCategoryMarkupRateDetailAPI Agent Markup Category Markup Rate Detail Api
type AgentMarkupCategoryMarkupRateDetailAPI struct {
	Markup                     *AgentMarkupRateAPI      `json:"markup" gorm:"-"`
	MarkupRateDestinationGroup *[]MRDestinationGroupAPI `json:"markup_rate_destination_group" validate:"required,gte=1" gorm:"-"`
	MarkupRateAirlineCategory  *[]MRAirlineCategoryAPI  `json:"markup_rate_airline_category" gorm:"-"`
	MarkupRateAirline          *[]MRAirlineAPI          `json:"markup_rate_airline" gorm:"-"`
	MarkupRateSupplierType     *[]MRSupplierTypeAPI     `json:"markup_rate_supplier_type" gorm:"-"`
	MarkupRateHotelSupplier    *[]MRHotelSupplierAPI    `json:"markup_rate_hotel_supplier" gorm:"-"`
}

// MRDestinationGroupAPI Mr Destination Group Api
type MRDestinationGroupAPI struct {
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid" validate:"required,eq=16"`
	IsIncluded         *bool      `json:"is_included,omitempty"`
}

// MRAirlineCategoryAPI Mr Airline Category Api
type MRAirlineCategoryAPI struct {
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid" validate:"required,eq=16"`
	IsIncluded        *bool      `json:"is_included,omitempty"`
}

// MRAirlineAPI Mr Airline Api
type MRAirlineAPI struct {
	AirlineID  *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid" validate:"required,eq=16"`
	IsIncluded *bool      `json:"is_included,omitempty"`
}

// MRSupplierTypeAPI Mr Supplier Type Api
type MRSupplierTypeAPI struct {
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid" validate:"required,eq=16"`
	IsIncluded     *bool      `json:"is_included,omitempty"`
}

// MRHotelSupplierAPI Mr Hotel Supplier Api
type MRHotelSupplierAPI struct {
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid" validate:"required,eq=16"`
	IsIncluded      *bool      `json:"is_included,omitempty"`
}
