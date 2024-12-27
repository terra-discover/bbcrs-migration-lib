package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ServiceFeeRateProductType Model
type ServiceFeeRateProductType struct {
	basic.Base
	basic.DataOwner
	ServiceFeeRateProductTypeAPI
	ServiceFeeRate *ServiceFeeRate `json:"service_fee_rate" gorm:"foreignKey:ServiceFeeRateID;references:ID"`
	ProductType    *ProductType    `json:"product_type" gorm:"foreignKey:ProductTypeID;references:ID"`
}

// ServiceFeeRateProductTypeAPI API
type ServiceFeeRateProductTypeAPI struct {
	ServiceFeeRateID *uuid.UUID `json:"service_fee_rate_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_product_type__product_type_id,unique,where:deleted_at is null;not null;" format:"uuid"`
	ProductTypeID    *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);index:service_fee_rate_product_type__product_type_id,unique,where:deleted_at is null;not null;" format:"uuid"`
}
