package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SiteTypeTranslation SiteType Translation
type SiteTypeTranslation struct {
	basic.Base
	basic.DataOwner
	SiteTypeTranslationAPI
	SiteTypeID *uuid.UUID `json:"site_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:site_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Site Type id
}

// SiteTypeTranslationAPI SiteType Translation API
type SiteTypeTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:site_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SiteTypeName *string `json:"site_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Site Type Name
}
