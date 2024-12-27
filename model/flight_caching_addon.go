package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"
)

type FlightCachingAddon struct {
	basic.Base
	basic.DataOwner
	SessionID           *uuid.UUID                    `json:"session_id,omitempty" gorm:"type:varchar(36);index:idx_flight_caching_addon_session_id;not null" example:"0a739fb1-d80b-4514-b519-c75d422b1efc"`
	ParentID            *uuid.UUID                    `json:"parent_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggerignore:"true"` // ParentID
	PersonID            *uuid.UUID                    `json:"person_id,omitempty" gorm:"type:varchar(36)" format:"uuid" swaggerignore:"true"` // PersonID
	FareID              *uuid.UUID                    `json:"fare_id,omitempty" gorm:"type:varchar(36)"`
	RouteID             *uuid.UUID                    `json:"route_id,omitempty" gorm:"type:varchar(36)"`
	AddonType           *string                       `json:"addon_type,omitempty" gorm:"type:varchar(30);not null" example:"baggage"` //baggage/meal/seat/infant/drink/entertainment/sport/etc...
	Code                *string                       `json:"code,omitempty" example:"PBAB"`                                           //platform specific code
	ProductCategoryID   *uuid.UUID                    `json:"product_category_id,omitempty"`
	Title               *string                       `json:"title,omitempty" example:"Extra Baggage 20kg"`
	Description         *string                       `json:"description,omitempty" gorm:"type:text" example:"Lorem ipsum dolor sit amet"`
	CurrencyCode        *string                       `json:"currency_code,omitempty" gorm:"type:varchar(3)"` // Currency Code
	Price               *float64                      `json:"price,omitempty" example:"500000.00"`
	DiscountAmount      *float64                      `json:"discount_amount,omitempty" example:"500000.00"`
	TaxAmount           *float64                      `json:"tax_amount,omitempty" example:"500000.00"`
	PriceAfterMarkup    *float64                      `json:"price_after_markup,omitempty"` // will be filled by pricemodifier package
	Weight              *float64                      `json:"weight,omitempty" example:"7.5"`
	AddonBundle         *[]FlightCachingAddon         `json:"addon_bundle,omitempty" gorm:"foreignKey:ParentID;references:ID"`
	Assignees           *[]FlightCachingAddonAssignee `json:"assignees,omitempty"`
	FlightCachingRoutes *FlightCachingRoutes          `json:"route,omitempty" gorm:"foreignKey:RouteID"`
	ProductCategory     *ProductCategory              `json:"product_category,omitempty"`
	Additional          *string                       `json:"additional,omitempty"`
	IsNew               *bool                         `json:"is_new,omitempty" default:"false"`       // Difference New Addon Or Old Addon when modify addon
	IsRepricing         *bool                         `json:"is_repricing,omitempty" default:"false"` // Distinguish, and presented prices are the result of repricing
}

// GetFlightCachingAddon by query filter
func (data *FlightCachingAddon) GetFlightCachingAddon(tx *gorm.DB, queryFilters string) {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Preload(`AddonBundle`).Preload(`FlightCachingRoutes`).Preload(`ProductCategory`).Take(&data)
}

// GetFlightCachingAddons by query filter
func (data *FlightCachingAddon) GetFlightCachingAddons(tx *gorm.DB, queryFilters string) []FlightCachingAddon {
	res := []FlightCachingAddon{}
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).Preload(`AddonBundle`).Preload(`FlightCachingRoutes`).Preload(`ProductCategory`).Find(&res)
	return res
}

// BAGGAGE_QTY_MAP
var BAGGAGE_QTY_MAP = map[float64]string{
	10: "10 Kg",
	15: "15 Kg",
	20: "20 Kg",
	25: "25 Kg",
	30: "30 Kg",
	40: "40 Kg",
}
