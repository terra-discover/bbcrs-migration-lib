package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteTranslation Site Translation
type SiteTranslation struct {
	basic.Base
	basic.DataOwner
	SiteTranslationAPI
	SiteID *uuid.UUID `json:"site_id,omitempty" gorm:"type:varchar(36);uniqueIndex:site_translation_unique;not null" swaggertype:"string" format:"uuid"` // Site id
}

// SiteTranslationAPI Site Translation API
type SiteTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:site_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SiteName     *string `json:"site_name,omitempty" gorm:"type:varchar(64)" example:"Garuda Indonesia"`                                   // Site Name
}
