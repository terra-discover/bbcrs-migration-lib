package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteTheme SiteTheme
type SiteTheme struct {
	basic.Base
	basic.DataOwner
	SiteThemeAPI
	Site  *Site  `json:"site,omitempty" gorm:"foreignKey:SiteID;references:ID"`
	Theme *Theme `json:"theme,omitempty" gorm:"foreignKey:ThemeID;references:ID"`
}

// SiteThemeAPI SiteTheme API
type SiteThemeAPI struct {
	SiteID    *uuid.UUID       `json:"site_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	ThemeID   *uuid.UUID       `json:"theme_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	StartDate *strfmt.DateTime `json:"start_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
	EndDate   *strfmt.DateTime `json:"end_date,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"`
}
