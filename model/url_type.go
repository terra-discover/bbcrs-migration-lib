package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// URLType URL Type
type URLType struct {
	basic.Base
	basic.DataOwner
	URLTypeAPI
	URLTypeTranslation *URLTypeTranslation `json:"url_type_translation,omitempty"`
}

// URLTypeAPI URL Type API
type URLTypeAPI struct {
	URLTypeCode *int    `json:"url_type_code,omitempty" gorm:"type:smallint;not null;index:idx_url_type_code_deleted_at,unique,where:deleted_at is null" example:"1"`           // URL Type Code
	URLTypeName *string `json:"url_type_name,omitempty" gorm:"type:varchar(256);not null;index:idx_url_type_name_deleted_at,unique,where:deleted_at is null" example:"Website"` // URL Type Name
}

// Seed URLType
func (urlType *URLType) Seed() *URLType {
	seed := URLType{
		URLTypeAPI: URLTypeAPI{
			URLTypeCode: lib.Intptr(1),
			URLTypeName: lib.Strptr("Website"),
		},
	}
	_ = lib.Merge(seed, &urlType)
	return urlType
}

// URLTypes struct
type URLTypes []URLType

// Seed for seed
func (urlTypes *URLTypes) Seed() *URLTypes {
	name := []string{"Website", "Corporate Website", "Additional Website"}
	for i := 1; i < len(name); i++ {
		*urlTypes = append(*urlTypes, URLType{URLTypeAPI: URLTypeAPI{
			URLTypeCode: lib.Intptr(i + 1),
			URLTypeName: lib.Strptr(name[i]),
		}})
	}
	return urlTypes
}
