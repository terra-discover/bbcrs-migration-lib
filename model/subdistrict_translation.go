package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SubdistrictTranslation Subdistrict Translation
type SubdistrictTranslation struct {
	basic.Base
	basic.DataOwner
	SubdistrictTranslationAPI
	SubdistrictID *uuid.UUID `json:"subdistrict_id,omitempty" gorm:"type:varchar(36);uniqueIndex:subdistrict_translation_unique;not null" swaggertype:"string" format:"uuid"` // Subdistrict id
}

// SubdistrictTranslationAPI Subdistrict Translation API
type SubdistrictTranslationAPI struct {
	LanguageCode    *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:subdistrict_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SubdistrictName *string `json:"subdistrict_name,omitempty" gorm:"type:varchar(256)" example:"Gambir"`                                            // Subdistrict Name
}
