package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// DistrictTranslation District Translation
type DistrictTranslation struct {
	basic.Base
	basic.DataOwner
	DistrictTranslationAPI
	DistrictID *uuid.UUID `json:"district_id,omitempty" gorm:"type:varchar(36);uniqueIndex:district_translation_unique;not null" swaggertype:"string" format:"uuid"` // District id
}

// DistrictTranslationAPI District Translation API
type DistrictTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:district_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	DistrictName *string `json:"district_name,omitempty" gorm:"type:varchar(256)" example:"Gambir"`                                            // District Name
}
