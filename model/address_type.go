package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AddressType Address Type
type AddressType struct {
	basic.Base
	basic.DataOwner
	AddressTypeAPI
	AddressTypeTranslation *AddressTypeTranslation `json:"address_type_translation,omitempty"`
}

// AddressTypeAPI Address Type API
type AddressTypeAPI struct {
	AddressTypeCode *int    `json:"address_type_code,omitempty" gorm:"type:smallint;not null;index:idx_address_type_code_deleted_at,unique,where:deleted_at is null" example:"1"`            // Address Type Code
	AddressTypeName *string `json:"address_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_address_type_name_deleted_at,unique,where:deleted_at is null" example:"Delivery"` // Address Type Name
}

// Seed AddressType
func (addressType *AddressType) Seed() *AddressType {
	seed := AddressType{
		AddressTypeAPI: AddressTypeAPI{
			AddressTypeCode: lib.Intptr(1),
			AddressTypeName: lib.Strptr("Delivery"),
		},
	}
	_ = lib.Merge(seed, &addressType)
	return addressType
}

// AddressTypes slices
type AddressTypes []AddressType

// Seed func
func (addressTypes *AddressTypes) Seed() *AddressTypes {
	name := []string{"Delivery", "Mailing", "Billing", "Credit Card", "Other", "Contact", "Physical", "Pre-opening office", "Collection", "Chain", "Deposit", "Hotel", "Permanent", "Attraction"}
	j := 1
	for i := 0; i < len(name); i++ {
		if name[i] == "Attraction" {
			j = 1001
		}
		*addressTypes = append(*addressTypes, AddressType{
			AddressTypeAPI: AddressTypeAPI{
				AddressTypeCode: lib.Intptr(j),
				AddressTypeName: lib.Strptr(name[i]),
			},
		})
		j++
	}
	return addressTypes
}
