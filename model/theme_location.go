package model

import (
	"github.com/google/uuid"
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ThemeLocation Theme Location
type ThemeLocation struct {
	basic.Base
	basic.DataOwner
	ThemeLocationAPI
	Theme                    *Theme                    `json:"theme,omitempty"`
	ThemeLocationTranslation *ThemeLocationTranslation `json:"theme_location_translation,omitempty"`
}

// ThemeLocationAPI Theme Location API
type ThemeLocationAPI struct {
	ThemeLocationCode *string    `json:"theme_location_code,omitempty" gorm:"type:varchar(36);not null;index:idx_theme_location_code_deleted_at,unique,where:deleted_at is null" example:"main-menu"` // Theme Location Code
	ThemeLocationName *string    `json:"theme_location_name,omitempty" gorm:"type:varchar(256);not null;index:idx_theme_location_name_deleted_at,unique,where:deleted_at is null" example:""`         // Theme Location Name
	ThemeID           *uuid.UUID `json:"theme_id,omitempty" gorm:"type:varchar(36)"`                                                                                                                  // Theme ID
	IsAd              *bool      `json:"is_ad,omitempty" gorm:"default:false" example:"false"`                                                                                                        // Is Ad
	IsNews            *bool      `json:"is_news,omitempty" gorm:"default:false" example:"false"`                                                                                                      // Is News
	Description       *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                      // Description
}

// Seed ThemeLocation
func (themeLocation *ThemeLocation) Seed() *ThemeLocation {
	seed := ThemeLocation{
		ThemeLocationAPI: ThemeLocationAPI{
			ThemeLocationCode: lib.Strptr("main-menu"),
			ThemeLocationName: lib.Strptr("main"),
			ThemeID:           lib.UUIDPtr(uuid.New()),
		},
	}
	_ = lib.Merge(seed, &themeLocation)
	return themeLocation
}
