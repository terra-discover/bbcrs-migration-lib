package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RegionTranslation Region Translation
type RegionTranslation struct {
	basic.Base
	basic.DataOwner
	RegionTranslationAPI
	RegionID *uuid.UUID `json:"region_id,omitempty" gorm:"type:varchar(36);uniqueIndex:region_translation_unique;not null" swaggertype:"string" format:"uuid"` // Region id
}

// RegionTranslationAPI Region Translation API
type RegionTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:region_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RegionName   *string `json:"region_name,omitempty" gorm:"type:varchar(256)" example:"Indonesia"`                                         // Region Name
}
