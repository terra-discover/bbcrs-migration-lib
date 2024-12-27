package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// TaxRateProductType model
type TaxRateProductType struct {
	basic.Base
	basic.DataOwner
	TaxRateProductTypeAPI
	ProductType *ProductType `json:"product_type,omitempty"`
}

// TaxRateProductTypeAPI model
type TaxRateProductTypeAPI struct {
	TaxRateID     *uuid.UUID `json:"tax_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ProductTypeID *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
