package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MarkupRateSupplierType Model
type MarkupRateSupplierType struct {
	basic.Base
	basic.DataOwner
	MarkupRateSupplierTypeAPI
	MarkupRate   *MarkupRate   `json:"markup_rate" gorm:"foreignKey:MarkupRateID;references:ID"`
	SupplierType *SupplierType `json:"supplier_type" gorm:"foreignKey:SupplierTypeID;references:ID"`
}

// MarkupRateSupplierTypeAPI API
type MarkupRateSupplierTypeAPI struct {
	MarkupRateID   *uuid.UUID `json:"markup_rate_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	SupplierTypeID *uuid.UUID `json:"supplier_type_id,omitempty" gorm:"type:varchar(36);not null;" format:"uuid"`
	IsIncluded     *bool      `json:"is_included,omitempty"`
}
