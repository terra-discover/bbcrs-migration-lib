package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRate Model
type ServiceFeeRate struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateAPI
	ServiceFeeCategory             *ServiceFeeCategory              `json:"service_fee_category,omitempty"`
	FeeTaxType                     *FeeTaxType                      `json:"fee_tax_type,omitempty"`
	Currency                       *Currency                        `json:"currency,omitempty"`
	ChargeType                     *ChargeType                      `json:"charge_type,omitempty"`
	ParentServiceFeeRate           *ServiceFeeRate                  `json:"parent_service_fee_rate,omitempty" swaggerignore:"true"`
	ServiceFeeRateTranslation      *ServiceFeeRateTranslation       `json:"service_fee_rate_translation,omitempty"`
	ServiceFeeRateProductType      *ServiceFeeRateProductType       `json:"service_fee_rate_product_type,omitempty" swaggerignore:"true"`
	ServiceFeeRateDestinationGroup []ServiceFeeRateDestinationGroup `json:"service_fee_rate_destination_group,omitempty" gorm:"foreignKey:ServiceFeeRateID;references:ID" swaggerignore:"true"`
	ServiceFeeRateAirline          []ServiceFeeRateAirline          `json:"service_fee_rate_airline,omitempty" gorm:"foreignKey:ServiceFeeRateID;references:ID" swaggerignore:"true"`
	ServiceFeeRateAirlineCategory  []ServiceFeeRateAirlineCategory  `json:"service_fee_rate_airline_category,omitempty" gorm:"foreignKey:ServiceFeeRateID;references:ID" swaggerignore:"true"`
	ServiceFeeRateSupplierType     []ServiceFeeRateSupplierType     `json:"service_fee_rate_supplier_type,omitempty" gorm:"foreignKey:ServiceFeeRateID;references:ID" swaggerignore:"true"`
	ServiceFeeRateHotelSupplier    []ServiceFeeRateHotelSupplier    `json:"service_fee_rate_hotel_supplier,omitempty" gorm:"foreignKey:ServiceFeeRateID;references:ID" swaggerignore:"true"`
}

// ServiceFeeRateAPI API
type ServiceFeeRateAPI struct {
	ServiceFeeCategoryID       *uuid.UUID `json:"service_fee_category_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	FeeTaxTypeID               *uuid.UUID `json:"fee_tax_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	CurrencyID                 *uuid.UUID `json:"currency_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ChargeTypeID               *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	ParentServiceFeeRateID     *uuid.UUID `json:"parent_service_fee_rate_id,omitempty" gorm:"type:varchar(36);" format:"uuid"`
	IsSummary                  *bool      `json:"is_summary,omitempty"`
	Percent                    *float64   `json:"percent,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`
	Amount                     *float64   `json:"amount,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`
	IsTaxInclusive             *bool      `json:"is_tax_inclusive,omitempty"`
	Cascade                    *bool      `json:"cascade,omitempty"`
	ApplicationOrder           *int64     `json:"application_order,omitempty" example:"10"`
	AssignAllProductTypes      *bool      `json:"assign_all_product_types,omitempty" gorm:"default:false"`
	AssignAllProductCategories *bool      `json:"assign_all_product_categories,omitempty" gorm:"default:false"`
	AssignAllDestinationGroups *bool      `json:"assign_all_destination_groups,omitempty" gorm:"default:false"`
	AssignAllAirlineCategories *bool      `json:"assign_all_airline_categories,omitempty" gorm:"default:false"`
	AssignAllAirlines          *bool      `json:"assign_all_airlines,omitempty" gorm:"default:false"`
	AssignAllSupplierTypes     *bool      `json:"assign_all_supplier_types,omitempty" gorm:"default:false"`
	AssignAllHotelSuppliers    *bool      `json:"assign_all_hotel_suppliers,omitempty" gorm:"default:false"`
	Description                *string    `json:"description,omitempty" gorm:"type:text"`
}
