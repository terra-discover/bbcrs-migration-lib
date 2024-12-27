package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// ProductType Product Type
type ProductType struct {
	basic.Base
	basic.DataOwner
	ProductTypeAPI
	ProductTypeTranslation *ProductTypeTranslation `json:"product_type_translation,omitempty"`
}

// ProductTypeAPI Product Type API
type ProductTypeAPI struct {
	ProductTypeCode     *int       `json:"product_type_code,omitempty" gorm:"type:smallint;not null;index:idx_product_type_code_deleted_at,unique,where:deleted_at is null" example:"1"`          // Product Type Code
	ProductTypeName     *string    `json:"product_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_product_type_name_deleted_at,unique,where:deleted_at is null" example:"Flight"` // Product Type Name
	ParentProductTypeID *uuid.UUID `json:"parent_product_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid" example:"70fb0d70-05dd-46e2-b113-b56437a8b694"`            // Parent Product Type ID
	IsDefault           *bool      `json:"is_default,omitempty" example:"true"`                                                                                                                   // Is Default
}

// Seed data
func (ProductType) Seed() *[]ProductType {
	data := []ProductType{}
	items := []string{
		"Flight",
		"Hotel",
		"Other",
		"Tour",
		"Document",
	}

	for i := range items {
		var code int = i + 1
		var name string = items[i]
		data = append(data, ProductType{
			ProductTypeAPI: ProductTypeAPI{
				ProductTypeCode: &code,
				ProductTypeName: &name,
				IsDefault:       lib.Boolptr(i == 0),
			},
		})
	}

	return &data
}

// SeedOne ProductType
func (productType *ProductType) SeedOne() *ProductType {
	seed := ProductType{
		ProductTypeAPI: ProductTypeAPI{
			ProductTypeCode: lib.Intptr(1),
			ProductTypeName: lib.Strptr("Hotel"),
		},
	}
	_ = lib.Merge(seed, &productType)
	return productType
}

// GetProductType by query filter
func (data *ProductType) GetProductType(tx *gorm.DB, queryFilters string) *ProductType {
	qf, wf, _, _ := lib.CustomFilters(queryFilters, "", "")
	tx.Where(qf, wf...).First(&data)
	return data
}
