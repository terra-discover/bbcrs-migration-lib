package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

// ProductCategory Model
type ProductCategory struct {
	basic.Base
	basic.DataOwner
	ProductCategoryAPI
	ParentProductCategory *ProductCategory `json:"product_category,omitempty" swaggerignore:"true"`
	// TODO: prepare FlightCachingAddon into one package
	FlightCachingAddon []FlightCachingAddon `json:"flight_caching_addon,omitempty"`
}

// ProductCategoryAPI Model
type ProductCategoryAPI struct {
	ProductCategoryCode     *string    `json:"product_category_code,omitempty" gorm:"type:varchar(36);not null;index:idx_product_category_code_deleted_at,unique,where:deleted_at is null" example:"seat"` // Product Category Code
	ProductCategoryName     *string    `json:"product_category_name,omitempty" gorm:"type:varchar(256);not null"`                                                                                          // Product Category Name
	ParentProductCategoryID *uuid.UUID `json:"parent_product_category,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                               // Parent Product Category
	IsMandatory             *bool      `json:"is_mandatory,omitempty"`                                                                                                                                     // Is Mandatory
	IsDefault               *bool      `json:"is_default,omitempty"`                                                                                                                                       // Is Default
	Description             *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                     // Description
}

// GetProductCategory by query filter
func (data *ProductCategory) GetProductCategory(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&data)
}

// Seed data
func (s ProductCategory) Seed() *[]ProductCategory {
	data := []ProductCategory{}
	items := []string{
		"seat|Seat",
		"baggage|Baggage",
		"seatenhancement|Seat Enhancement (Leg Room, Forward Seat)",
		"meal|Meal",
		"drink|Drink",
		"inflightentertainment|In-Flight Entertainment",
		"sportequipment|Sport Equipment",
		"infant|Infant Equipment",
		"flightother|Flight Other Product",
		"travelinsurance|Travel Insurance",
		"visa|Visa",
		"passport|Passport",
		"rentcar|Rent Car",
		"other|Other Product",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		code := c[0]
		name := c[1]

		data = append(data, ProductCategory{
			ProductCategoryAPI: ProductCategoryAPI{
				ProductCategoryCode: &code,
				ProductCategoryName: &name,
			},
		})
	}

	return &data
}
