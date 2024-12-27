package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteLanguage SiteLanguage
type SiteLanguage struct {
	basic.Base
	basic.DataOwner
	SiteLanguageAPI
	Site     *Site     `json:"site,omitempty" gorm:"foreignKey:SiteID;references:ID"`
	Language *Language `json:"language,omitempty" gorm:"foreignKey:LanguageID;references:ID"`
}

// SiteLanguageAPI SiteLanguage API
type SiteLanguageAPI struct {
	SiteID     *uuid.UUID `json:"site_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	LanguageID *uuid.UUID `json:"language_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);not null;"`
	IsPrimary  *bool      `json:"is_primary,omitempty"`
}
