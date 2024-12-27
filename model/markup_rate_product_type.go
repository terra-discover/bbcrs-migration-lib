package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateProductType Model
type MarkupRateProductType struct {
	basic.Base
	basic.DataOwner
	MarkupRateProductTypeAPI
	MarkupRate  *MarkupRate  `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	ProductType *ProductType `json:"product_type" gorm:"foreignKey:ProductTypeID;references:ID"`
}

// MarkupRateProductTypeAPI API
type MarkupRateProductTypeAPI struct {
	MarkupRateID  *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	ProductTypeID *uuid.UUID `json:"product_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
}
