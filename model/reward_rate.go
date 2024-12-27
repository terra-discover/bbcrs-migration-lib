package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RewardRate Reward Rate
type RewardRate struct {
	basic.Base
	basic.DataOwner
	RewardRateAPI
}

// RewardRateAPI Reward Rate API
type RewardRateAPI struct {
	RewardCategoryID           *uuid.UUID `json:"reward_category_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"` // Reward Category Id
	Percent                    *float32   `json:"percent,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`                                        // Percent
	RewardTypeID               *uuid.UUID `json:"reward_type_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`     // Reward Type Id
	ChargeTypeID               *uuid.UUID `json:"charge_type_id,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"`     // Reward Type Id
	Amount                     *float32   `json:"amount,omitempty" gorm:"type:decimal(19,4)" example:"1000.00"`                                         // Amount
	IsTaxInclusive             *bool      `json:"is_tax_inclusive,omitempty"`                                                                           // Is Tax Inclusive
	AssignAllProductTypes      *bool      `json:"assign_all_product_types,omitempty"`                                                                   // Assign All Product Types
	AssignAllProductCategories *bool      `json:"assign_all_product_categories,omitempty"`                                                              // Assign All Product Categories
	AssignAllDestinationGroups *bool      `json:"assign_all_destination_groups,omitempty"`                                                              // Assign All Destination Groups
	AssignAllAirlineCategories *bool      `json:"assign_all_airline_categories,omitempty"`                                                              // Assign All Airline Categories
	AssignAllAirline           *bool      `json:"assign_all_airlines,omitempty"`                                                                        // Assign All Airline
	AssignAllSupplierTypes     *bool      `json:"assign_all_supplier_types,omitempty"`                                                                  // Assign All Supplier Types
	AssignAllHotelSuppliers    *bool      `json:"assign_all_hotel_suppliers,omitempty"`                                                                 // Assign All Hotel Suppliers
	Description                *string    `json:"description,omitempty" gorm:"type:varchar(4000)"`                                                      // Description
}
