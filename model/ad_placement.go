package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdPlacement Ad Placement
type AdPlacement struct {
	basic.Base
	basic.DataOwner
	AdPlacementAPI
	AdPlacementTranslation *AdPlacementTranslation `json:"ad_placement_translation,omitempty"`
}

// AdPlacementAPI Ad Placement API
type AdPlacementAPI struct {
	AdPlacementCode *string `json:"ad_placement_code,omitempty" gorm:"type:varchar(36);not null;index:,unique,where:deleted_at is null"` // Ad Placement Code
	AdPlacementName *string `json:"ad_placement_name,omitempty" gorm:"type:varchar(64);not null;index:,unique,where:deleted_at is null"` // Ad Placement Name
	Comment         *string `json:"comment,omitempty" gorm:"type:text"`                                                                  // Comment
	Description     *string `json:"description,omitempty" gorm:"type:text"`                                                              // Description
}

// Seed AdPlacement
func (adPlacement *AdPlacement) Seed() *AdPlacement {
	seed := AdPlacement{
		AdPlacementAPI: AdPlacementAPI{
			AdPlacementCode: lib.Strptr("Ad Placement Code"),
			AdPlacementName: lib.Strptr("Ad Placement Name"),
			Comment:         nil,
			Description:     nil,
		},
	}
	_ = lib.Merge(seed, &adPlacement)
	return adPlacement
}
