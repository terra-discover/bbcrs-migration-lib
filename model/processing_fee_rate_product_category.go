package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateProductCategory Model
type ProcessingFeeRateProductCategory struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateProductCategoryAPI
	ProcessingFeeRate *ProcessingFeeRate `json:"processing_fee_rate,omitempty"` // ProcessingFee Rate
	ProductCategory   *ProductCategory   `json:"product_category,omitempty"`    // Product Category
}

// ProcessingFeeRateProductCategoryAPI Model
type ProcessingFeeRateProductCategoryAPI struct {
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // ProcessingFee Rate ID
	ProductCategoryID   *uuid.UUID `json:"product_category_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`    // Product Category ID
}
