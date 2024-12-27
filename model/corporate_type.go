package model

import (
	"strings"

	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

	"github.com/google/uuid"
)

// CorporateType Corporate Type
type CorporateType struct {
	basic.Base
	basic.DataOwner
	CorporateTypeAPI
}

// CorporateTypeAPI Corporate Type API
type CorporateTypeAPI struct {
	CorporateTypeCode *string `json:"corporate_type_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null" swaggertype:"string" format:"uuid"` // Corporate Type Code
	CorporateTypeName *string `json:"corporate_type_name,omitempty" gorm:"type:varchar(256);not null"`                                                                          // Corporate Type Name
	IsDefault         *bool   `json:"is_default,omitempty"`                                                                                                                     // Is Default
}

// Seed Site
func (p CorporateType) Seed() *[]CorporateType {
	data := []CorporateType{}
	items := []string{
		"1b7d2d4b-8934-447f-8812-19d579cdb7b2|self|Self Service",
		"c5f84e5f-b75c-45d2-be87-1bdb7c160b76|shared|Shared",
		"1cdc048e-3373-4736-81f3-a4d7e421139e|dedicated|Dedicated",
		"b81d9769-705b-46f7-b625-777655c01394|inplant|Dedicated In Plant",
	}

	for i := range items {
		values := strings.Split(items[i], "|")

		id := uuid.MustParse(values[0])
		code := values[1]
		name := values[2]
		isDefault := false

		// Must set 1 record with is_default to true
		if code == "self" {
			isDefault = true
		}

		data = append(data, CorporateType{
			Base: basic.Base{
				ID: &id,
			},
			CorporateTypeAPI: CorporateTypeAPI{
				CorporateTypeCode: &code,
				CorporateTypeName: &name,
				IsDefault:         lib.Boolptr(isDefault),
			},
		})
	}

	return &data
}
