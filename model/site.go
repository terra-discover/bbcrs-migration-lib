package model

import (
	"strings"

	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Site Site
type Site struct {
	basic.Base
	basic.DataOwner
	SiteAPI
	SiteType        *SiteType        `json:"site_type,omitempty"`     // Site Type
	TimeZone        *TimeZone        `json:"time_zone,omitempty"`     // Timezone
	BuiltInType     *BuiltInType     `json:"built_in_type,omitempty"` // BuiltInType
	SiteAsset       *SiteAsset       `json:"site_asset,omitempty"`
	SiteTranslation *SiteTranslation `json:"site_translation,omitempty"`
}

// SiteAPI Site API
type SiteAPI struct {
	SiteCode      *string       `json:"site_code,omitempty" gorm:"type:varchar(36);not null;index:idx_site_code_deleted_at,unique,where:deleted_at is null" validate:"required,gt=0,lte=36"`   // Site Code
	SiteName      *string       `json:"site_name,omitempty" gorm:"type:varchar(256);not null;index:idx_site_name_deleted_at,unique,where:deleted_at is null" validate:"required,gt=0,lte=256"` // Site Name
	BuiltInTypeID *uuid.UUID    `json:"built_in_type_id,omitempty" swaggertype:"string" gorm:"type:varchar(36);not null;" format:"uuid"`
	SiteTypeID    *uuid.UUID    `json:"site_type_id,omitempty" swaggertype:"string" gorm:"type:varchar(36);not null;" format:"uuid"`
	TimeZoneID    *uuid.UUID    `json:"time_zone_id,omitempty" swaggertype:"string" gorm:"type:varchar(36);" format:"uuid"`
	SiteAsset     *SiteAssetAPI `json:"site_asset,omitempty" gorm:"-"`
}

// Seed Site
// func (site *Site) Seed() *Site {
// 	seed := Site{
// 		SiteAPI: SiteAPI{
// 			SiteCode:      lib.Strptr("Code"),
// 			SiteName:      lib.Strptr("Site 1"),
// 			SiteTypeID:    lib.UUIDPtr(uuid.New()),
// 			TimeZoneID:    lib.UUIDPtr(uuid.New()),
// 			BuiltInTypeID: lib.UUIDPtr(uuid.New()),
// 		},
// 	}
// 	_ = lib.Merge(seed, &site)
// 	return site
// }

// Seed Site
func (s Site) Seed() *[]Site {
	data := []Site{}
	items := []string{
		"Code|Site 1|a29e91e4-a1b6-4612-9ce6-25c4bc96a745|cb0a1c2b-b1a8-4571-927b-acff3824dcd8|f9c7b22b-2494-4500-865a-910e32e6ab39",
		"Code 1|Site 2|5cddf574-b6db-492b-8288-8a04386436f2|cb0a1c2b-b1a8-4571-927b-acff3824dcd8|07886990-c7eb-42a5-b49b-8b982644e247",
		"Code 2|Site 3|bedb5f42-6d8a-4026-a1ac-ca8a6704f469|cb0a1c2b-b1a8-4571-927b-acff3824dcd8|866bc058-8623-4997-ab15-26567ed8b623",
	}

	for i := range items {
		values := strings.Split(items[i], "|")
		var buildInTypeID uuid.UUID
		var siteTypeID uuid.UUID
		var timeZoneID uuid.UUID

		buildInTypeID, _ = uuid.Parse(values[2])
		var code string = values[0]
		var name string = values[1]
		siteTypeID, _ = uuid.Parse(values[3])
		timeZoneID, _ = uuid.Parse(values[4])
		data = append(data, Site{
			SiteAPI: SiteAPI{
				SiteCode:      &code,
				SiteName:      &name,
				BuiltInTypeID: &buildInTypeID,
				SiteTypeID:    &siteTypeID,
				TimeZoneID:    &timeZoneID,
			},
		})
	}

	return &data
}
