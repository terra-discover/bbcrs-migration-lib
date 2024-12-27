package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AgentServiceFeeCategory Agent Service Fee Category
type AgentServiceFeeCategory struct {
	basic.Base
	basic.DataOwner
	AgentServiceFeeCategoryAPI
	ServiceFeeCategory *ServiceFeeCategory `json:"service_fee_category"`
	Agent              *Agent              `json:"agent" gorm:"foreignKey:AgentID;references:ID"`
}

// AgentServiceFeeCategoryAPI Agent Service Fee Category API
type AgentServiceFeeCategoryAPI struct {
	AgentID              *uuid.UUID `json:"agent_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`                // Agent ID
	ServiceFeeCategoryID *uuid.UUID `json:"service_fee_category_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"` // Service Fee Category ID
}

// AgentServiceFeeCategoryDetail struct
type AgentServiceFeeCategoryDetail struct {
	basic.Base
	basic.DataOwner
	AgentServiceFeeCategoryDetailAPI
	NumberOfOverride               *int64          `json:"number_of_override,omitempty"`
	DomesticFlightServiceFee       *ServiceFeeRate `json:"domestic_flight_service_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalFlightServiceFee  *ServiceFeeRate `json:"international_flight_service_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticHotelServiceFee        *ServiceFeeRate `json:"domestic_hotel_service_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalHotelServiceFee   *ServiceFeeRate `json:"international_hotel_service_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticFlightServiceFees      *string         `json:"domestic_flight_service_fees,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalFlightServiceFees *string         `json:"international_flight_service_fees,omitempty" gorm:"foreignKey:ID;references:ID"`
	DomesticHotelServiceFees       *string         `json:"domestic_hotel_service_fees,omitempty" gorm:"foreignKey:ID;references:ID"`
	InternationalHotelServiceFees  *string         `json:"international_hotel_service_fees,omitempty" gorm:"foreignKey:ID;references:ID"`
}

type AgentServiceFeeCategoryDetailAPI struct {
	ServiceFeeCategoryName        *string              `json:"service_fee_category_name" example:"Service Fee Category 1" validate:"required"` // Service Fee Category Name
	Description                   *string              `json:"description,omitempty"`                                                          // Description
	DomesticFlightServiceFee      *AgentServiceFeeRate `json:"domestic_flight,omitempty" gorm:"-"`                                             // Domestic Service Fee Rate
	InternationalFlightServiceFee *AgentServiceFeeRate `json:"international_flight,omitempty" gorm:"-"`
	DomesticHotelServiceFee       *AgentServiceFeeRate `json:"domestic_hotel,omitempty" gorm:"-"`      // Domestic Service Fee Rate
	InternationalHotelServiceFee  *AgentServiceFeeRate `json:"international_hotel,omitempty" gorm:"-"` // International Service Fee Rate
}

type AgentServiceFeeRate struct {
	ChargeTypeID   *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" format:"uuid"` // Charge Type Id
	Percent        *float64   `json:"percent,omitempty" gorm:"type:decimal(19,4)"`                    // Percent
	Amount         *float64   `json:"amount,omitempty" gorm:"type:decimal(19,4)"`                     // Amount
	IsTaxInclusive *bool      `json:"is_tax_inclusive,omitempty"`                                     // Is Tax Inclusive
}

type AgentServiceFeeCategoryServiceFeeRateDetail struct {
	basic.Base
	basic.DataOwner
	ServiceFee            *ServiceFeeRate   `json:"service_fee_rate,omitempty" gorm:"foreignKey:ID;references:ID"`
	DestinationGroup      *DestinationGroup `json:"destination_group,omitempty" gorm:"foreignKey:ID"`
	AirlineCategory       *AirlineCategory  `json:"airline_category,omitempty" gorm:"foreignKey:ID"`
	Airline               *Airline          `json:"airline,omitempty" gorm:"foreignKey:ID"`
	DestinationGroupNames *string           `json:"destination_group_names,omitempty"` // Destination Group Names joined by ampersand `,`
	AirlineNames          *string           `json:"airline_names,omitempty"`           // Airline Names joined by ampersand `,`
	AirlineCategoryNames  *string           `json:"airline_category_names,omitempty"`  // Airline Category Names joined by ampersand `,`
	SupplierTypeNames     *string           `json:"supplier_type_names,omitempty"`     // Supplier Type Names joined by ampersand `,`
	HotelSupplierNames    *string           `json:"hotel_supplier_names,omitempty"`    // Hotel Supplier Names joined by ampersand `,`
	ApplyCondition        *string           `json:"apply_condition,omitempty"`         // Apply Condition
	ServiceFees           *string           `json:"service_fees,omitempty"`            // Service Fees joined by ampersand `,`
}

type AgentServiceFeeCategoryServiceFeeRateResultDetail struct {
	basic.Base
	basic.DataOwner
	AgentServiceFeeCategoryServiceFeeRateDetailAPI
	ServiceFee *ServiceFeeRate `json:"service_fee,omitempty" gorm:"foreignKey:ID;references:ID"`
}

type AgentServiceFeeCategoryServiceFeeRateDetailAPI struct {
	ServiceFee                     *AgentServiceFeeRate      `json:"service_fee" gorm:"-"`
	ServiceFeeRateDestinationGroup *[]SFRDestinationGroupAPI `json:"service_fee_rate_destination_group" gorm:"-"`
	ServiceFeeRateAirlineCategory  *[]SFRAirlineCategoryAPI  `json:"service_fee_rate_airline_category" gorm:"-"`
	ServiceFeeRateAirline          *[]SFRAirlineAPI          `json:"service_fee_rate_airline" gorm:"-"`
	ServiceFeeRateSupplierType     *[]SFRSupplierTypeAPI     `json:"service_fee_rate_supplier_type" gorm:"-"`
	ServiceFeeRateHotelSupplier    *[]SFRHotelSupplierAPI    `json:"service_fee_rate_hotel_supplier" gorm:"-"`
}

type SFRDestinationGroupAPI struct {
	DestinationGroupID *uuid.UUID `json:"destination_group_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded         *bool      `json:"is_included,omitempty"`
}

type SFRAirlineCategoryAPI struct {
	AirlineCategoryID *uuid.UUID `json:"airline_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded        *bool      `json:"is_included,omitempty"`
}

type SFRAirlineAPI struct {
	AirlineID  *uuid.UUID `json:"airline_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded *bool      `json:"is_included,omitempty"`
}

type SFRSupplierTypeAPI struct {
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded     *bool      `json:"is_included,omitempty"`
}

type SFRHotelSupplierAPI struct {
	HotelSupplierID *uuid.UUID `json:"hotel_supplier_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded      *bool      `json:"is_included,omitempty"`
}

// FlightStandardServiceFee struct
type FlightStandardServiceFee struct {
	Percent        *float64   `json:"percent,omitempty" gorm:"-" example:"1000.00"`
	CurrencyID     *uuid.UUID `json:"currency_id,omitempty" gorm:"-" format:"uuid"`
	Amount         *float64   `json:"amount,omitempty" gorm:"-" example:"1000.00"`
	ChargeType     *string    `json:"charge_type,omitempty" gorm:"-" format:"uuid"`
	IsTaxInclusive *bool      `json:"is_tax_inclusive,omitempty" gorm:"-"`
	FeeTaxTypeID   *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"-" format:"uuid"`
	FlagType       *string    `json:"flag_type,omitempty" gorm:"-"  example:"amount or percent"`
}
