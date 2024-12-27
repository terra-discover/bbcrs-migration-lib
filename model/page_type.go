package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// PageType Page Type
type PageType struct {
	basic.Base
	basic.DataOwner
	PageTypeAPI
	PageTypeTranslation *PageTypeTranslation `json:"page_type_translation,omitempty"`
}

// PageTypeAPI Page Type API
type PageTypeAPI struct {
	PageTypeCode *int    `json:"page_type_code,omitempty" gorm:"type:smallint;not null;index:,unique,where:deleted_at is null" example:"1"`            // Page Type Code
	PageTypeName *string `json:"page_type_name,omitempty" gorm:"type:varchar(256);not null;index:,unique,where:deleted_at is null" example:"Feedback"` // Page Type Name
}

// Seed PageType
func (pageType *PageType) Seed() *PageType {
	seed := PageType{
		PageTypeAPI: PageTypeAPI{
			PageTypeCode: lib.Intptr(1),
			PageTypeName: lib.Strptr("Feedback"),
		},
	}
	_ = lib.Merge(seed, &pageType)
	return pageType
}

// PageTypes type
type PageTypes []PageType

// Seed for PageTypes
func (pageTypes *PageTypes) Seed() *PageTypes {
	name := []string{"General Inquiry", "Feedback", "Help", "Help Article", "How To Article", "Video Page", "News", "Destination Info"}
	for i := 0; i < len(name); i++ {
		*pageTypes = append(*pageTypes, PageType{PageTypeAPI: PageTypeAPI{
			PageTypeCode: lib.Intptr(i + 1),
			PageTypeName: lib.Strptr(name[i]),
		}})
	}
	return pageTypes
}
