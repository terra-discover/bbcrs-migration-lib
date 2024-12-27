package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SupplierType Supplier Type
type SupplierType struct {
	basic.Base
	basic.DataOwner
	SupplierTypeAPI
	SupplierTypeTranslation *SupplierTypeTranslation `json:"supplier_type_translation,omitempty"`
}

// SupplierTypeAPI Supplier Type API
type SupplierTypeAPI struct {
	SupplierTypeCode *string `json:"supplier_type_code,omitempty" gorm:"type:varchar(36);not null;index:idx_supplier_type_code_deleted_at,unique,where:deleted_at is null" example:"DIRECT"`  // Supplier Type Code
	SupplierTypeName *string `json:"supplier_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_supplier_type_name_deleted_at,unique,where:deleted_at is null" example:"Direct"` // Supplier Type Name
}

// Seed SupplierType
func (supplierType *SupplierType) Seed() *SupplierType {
	seed := SupplierType{
		SupplierTypeAPI: SupplierTypeAPI{
			SupplierTypeCode: lib.Strptr("DIRECT"),
			SupplierTypeName: lib.Strptr("Direct"),
		},
	}
	_ = lib.Merge(seed, &supplierType)
	return supplierType
}
