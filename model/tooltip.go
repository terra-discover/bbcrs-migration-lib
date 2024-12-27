package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// Tooltip Tooltip
type Tooltip struct {
	basic.Base
	basic.DataOwner
	TooltipAPI
	TooltipAsset       *TooltipAsset       `json:"tooltip_asset,omitempty"`
	TooltipTranslation *TooltipTranslation `json:"tooltip_translation,omitempty"`
}

// TooltipAPI Tooltip API
type TooltipAPI struct {
	TooltipCode  *string          `json:"tooltip_code,omitempty" gorm:"type:varchar(36);not null;index:idx_tooltip_code_deleted_at,unique,where:deleted_at is null" validate:"required,gt=0,lte=36"`   // Tooltip Code
	TooltipName  *string          `json:"tooltip_name,omitempty" gorm:"type:varchar(256);not null;index:idx_tooltip_name_deleted_at,unique,where:deleted_at is null" validate:"required,gt=0,lte=256"` // Tooltip Name
	AutoClose    *bool            `json:"auto_close,omitempty" example:"true"`
	Body         *string          `json:"body,omitempty" gorm:"type:text" example:"body"`
	BodyFileName *string          `json:"body_file_name,omitempty" gorm:"type:varchar(256)"`
	Description  *string          `json:"description,omitempty" gorm:"type:text" example:"description"`
	TooltipAsset *TooltipAssetAPI `json:"tooltip_asset,omitempty" gorm:"-"`
}

// Seed Tooltip
func (tooltip *Tooltip) Seed() *Tooltip {
	seed := Tooltip{
		TooltipAPI: TooltipAPI{
			TooltipCode:  lib.Strptr("Code"),
			TooltipName:  lib.Strptr("Tooltip 1"),
			AutoClose:    lib.Boolptr(false),
			Body:         lib.Strptr("body"),
			BodyFileName: lib.Strptr("body file name"),
			Description:  lib.Strptr("description"),
		},
	}
	_ = lib.Merge(seed, &tooltip)
	return tooltip
}
