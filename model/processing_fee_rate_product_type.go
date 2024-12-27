package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateProductType schema
type ProcessingFeeRateProductType struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateProductTypeAPI
	ProcessingFeeRate *ProcessingFeeRate `json:"processing_fee_rate,omitempty"` // Processing Fee Rate
	ProductType       *ProductType       `json:"product_type,omitempty"`        // Product Type
}

// ProcessingFeeRateProductTypeAPI schema
type ProcessingFeeRateProductTypeAPI struct {
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);not null;" swaggertype:"string" format:"uuid"`
	ProductTypeID       *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
}
