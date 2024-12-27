package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateProductCategory Model
type MarkupRateProductCategory struct {
	basic.Base
	basic.DataOwner
	MarkupRateProductCategoryAPI
	MarkupRate      *MarkupRate      `json:"markup_rate,omitempty"`      // Markup Rate
	ProductCategory *ProductCategory `json:"product_category,omitempty"` // Product Category
}

// MarkupRateProductCategoryAPI Model
type MarkupRateProductCategoryAPI struct {
	MarkupRateID      *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`      // Markup Rate ID
	ProductCategoryID *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Product Category ID
}
