package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteType model
type SiteType struct {
	basic.Base
	basic.DataOwner
	SiteTypeAPI
	SiteTypeTranslation *SiteTypeTranslation `json:"site_type_translation,omitempty"`
}

// SiteTypeAPI model
type SiteTypeAPI struct {
	SiteTypeCode *string `json:"site_type_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`
	SiteTypeName *string `json:"site_type_name,omitempty" gorm:"type:varchar(256);not null"`
}

type SiteTypeCode string

const (
	BookingSiteTypeCode SiteTypeCode = "booking"
	CmsSiteTypeCode     SiteTypeCode = "cms"
	BackendSiteTypeCode SiteTypeCode = "backend"
)

// Seed SiteType
func (siteType *SiteType) Seed() *SiteType {
	seed := SiteType{
		Base: basic.Base{
			ID: lib.StringToUUID("cb0a1c2b-b1a8-4571-927b-acff3824dcd8"),
		},
		SiteTypeAPI: SiteTypeAPI{
			SiteTypeCode: lib.Strptr("booking"),
			SiteTypeName: lib.Strptr("Booking"),
		},
	}
	_ = lib.Merge(seed, &siteType)
	return siteType
}
