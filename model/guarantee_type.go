package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// GuaranteeType Guarantee Type
type GuaranteeType struct {
	basic.Base
	basic.DataOwner
	GuaranteeTypeAPI
	GuaranteeTypeTranslation *GuaranteeTypeTranslation `json:"guarantee_type_translation,omitempty"`
}

// GuaranteeTypeAPI Guarantee Type API
type GuaranteeTypeAPI struct {
	GuaranteeTypeCode *string `json:"guarantee_type_code,omitempty" gorm:"type:varchar(36);not null;index:idx_guarantee_type_code_deleted_at,unique,where:deleted_at is null" example:"GuaranteeRequired"`  // Guarantee Type Code
	GuaranteeTypeName *string `json:"guarantee_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_guarantee_type_name_deleted_at,unique,where:deleted_at is null" example:"GuaranteeRequired"` // Guarantee Type Name
}

// Seed GuaranteeType
func (guaranteeType *GuaranteeType) Seed() *GuaranteeType {
	seed := GuaranteeType{
		GuaranteeTypeAPI: GuaranteeTypeAPI{
			GuaranteeTypeCode: lib.Strptr("GuaranteeRequired"),
			GuaranteeTypeName: lib.Strptr("GuaranteeRequired"),
		},
	}
	_ = lib.Merge(seed, &guaranteeType)
	return guaranteeType
}
