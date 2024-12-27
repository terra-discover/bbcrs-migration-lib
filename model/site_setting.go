package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteSetting SiteSetting
type SiteSetting struct {
	basic.Base
	basic.DataOwner
	SiteSettingAPI
	Site *Site `json:"site,omitempty" gorm:"foreignKey:SiteID;references:ID"`
}

// SiteSettingAPI SiteSetting API
type SiteSettingAPI struct {
	SiteID *uuid.UUID `json:"site_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	Name   *string    `json:"name,omitempty" gorm:"type:varchar(256);not null;"`
	Value  *string    `json:"value,omitempty" gorm:"type:text"`
}
