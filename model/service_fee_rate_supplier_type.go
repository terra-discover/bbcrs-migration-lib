package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateSupplierType Model
type ServiceFeeRateSupplierType struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateSupplierTypeAPI
	ServiceFeeRate *ServiceFeeRate `json:"service_fee_rate" gorm:"foreignKey:ServiceFeeRateID;references:ID"`
	SupplierType   *SupplierType   `json:"supplier_type" gorm:"foreignKey:SupplierTypeID;references:ID"`
}

// ServiceFeeRateSupplierTypeAPI API
type ServiceFeeRateSupplierTypeAPI struct {
	ServiceFeeRateID *uuid.UUID `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_supplier_type__supplier_type_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	SupplierTypeID   *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_supplier_type__supplier_type_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	IsIncluded       *bool      `json:"is_included,omitempty" gorm:"default:true"`
}
