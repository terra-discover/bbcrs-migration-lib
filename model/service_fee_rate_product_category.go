package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateProductCategory Model
type ServiceFeeRateProductCategory struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateProductCategoryAPI
	ServiceFeeRate  *ServiceFeeRate  `json:"service_fee_rate,omitempty"` // ServiceFee Rate
	ProductCategory *ProductCategory `json:"product_category,omitempty"` // Product Category
}

// ServiceFeeRateProductCategoryAPI Model
type ServiceFeeRateProductCategoryAPI struct {
	ServiceFeeRateID  *uuid.UUID `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_product_category__product_category_id,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // ServiceFee Rate ID
	ProductCategoryID *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_product_category__product_category_id,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Product Category ID
}
