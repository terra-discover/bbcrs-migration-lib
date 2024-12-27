package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ProcessingFeeRateSupplierType Model
type ProcessingFeeRateSupplierType struct {
	basic.Base
	basic.DataOwner
	ProcessingFeeRateSupplierTypeAPI
	ProcessingFeeRate *ProcessingFeeRate `json:"processing_fee_rate" gorm:"foreignKey:ProcessingFeeRateID;references:ID"`
	SupplierType      *SupplierType      `json:"supplier_type" gorm:"foreignKey:SupplierTypeID;references:ID"`
}

// ProcessingFeeRateSupplierTypeAPI API
type ProcessingFeeRateSupplierTypeAPI struct {
	ProcessingFeeRateID *uuid.UUID `json:"processing_fee_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	SupplierTypeID      *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded          *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
