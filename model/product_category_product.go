package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Product Category Product
type ProductCategoryProduct struct {
	basic.Base
	basic.DataOwner
	ProductCategoryProductAPI
}

// ProductCategoryProductAPI Product Category Product API
type ProductCategoryProductAPI struct {
	ProductCategoryID *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	ProductID         *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
