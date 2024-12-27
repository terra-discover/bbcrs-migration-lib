package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// Product struct Product
type Product struct {
	basic.Base
	basic.DataOwner
	ProductAPI
	ProductTranslation *ProductTranslation `json:"product_translation,omitempty"`
	ProductAsset       *ProductAsset       `json:"product_asset,omitempty"`
	ProductType        *ProductType        `json:"product_type,omitempty"`
	Currency           *Currency           `json:"currency,omitempty"`
	//ChargeType         *ChargeType         `json:"charge_type,omitempty"`
}

// ProductAPI Product API
type ProductAPI struct {
	ProductCode            *string          `json:"product_code,omitempty" gorm:"type:varchar(36);not null;"`             // Product Code
	ProductName            *string          `json:"product_name,omitempty" gorm:"type:varchar(256);not null;"`            // Product Name
	ProductTypeID          *uuid.UUID       `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null;"`          // Product Type ID
	ProductUnit            *string          `json:"product_unit,omitempty" gorm:"type:varchar(32)"`                       // Product Unit
	ChargeTypeID           *uuid.UUID       `json:"charge_type_id,omitempty" gorm:"type:varchar(36)"`                     // Charge Type ID
	CurrencyID             *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)"`                        // Currency ID
	IsMandatory            *bool            `json:"is_mandatory,omitempty" example:"true"`                                // Is Mandatory
	IsTaxInclusive         *bool            `json:"is_tax_inclusive,omitempty" example:"true"`                            // Is Tax Inclusive
	TaxPercent             *float64         `json:"tax_percent,omitempty"`                                                // Tax Percent
	TaxAmount              *float64         `json:"tax_amount,omitempty"`                                                 // Tax Amount
	ManageInventory        *bool            `json:"manage_inventory,omitempty"`                                           // Manage Inventory
	Quantity               *int             `json:"quantity,omitempty" example:"1"`                                       // Quantity
	HasProductAvailability *bool            `json:"has_product_availability,omitempty" example:"true"`                    // Has Product Availability
	AssignAllRatePlans     *bool            `json:"assign_all_rate_plans,omitempty"`                                      // Assign All Rate Plans
	IsAccrual              *bool            `json:"is_accrual,omitempty" example:"true"`                                  // Is Accrual
	IsRedeemable           *bool            `json:"is_redeemable,omitempty" example:"true"`                               // Is Redeemable
	Description            *string          `json:"description,omitempty" gorm:"type:text" example:"description product"` // Description
	ProductAsset           *ProductAssetAPI `json:"product_asset,omitempty" gorm:"-"`                                     // Product Asset API
}

// Seed Product
func (product *Product) Seed() *Product {
	seed := Product{
		ProductAPI: ProductAPI{
			ProductCode:            lib.Strptr("Electronic"),
			ProductName:            lib.Strptr("Electronic"),
			ProductTypeID:          lib.UUIDPtr(uuid.New()),
			ProductUnit:            nil,
			ChargeTypeID:           nil,
			CurrencyID:             nil,
			IsMandatory:            nil,
			IsTaxInclusive:         nil,
			TaxPercent:             nil,
			TaxAmount:              nil,
			ManageInventory:        nil,
			Quantity:               nil,
			HasProductAvailability: nil,
			AssignAllRatePlans:     nil,
			IsAccrual:              nil,
			IsRedeemable:           nil,
			Description:            nil,
		},
	}
	_ = lib.Merge(seed, &product)
	return product
}
