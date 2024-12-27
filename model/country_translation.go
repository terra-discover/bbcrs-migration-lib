package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CountryTranslation struct is not writable by request body
type CountryTranslation struct {
	basic.Base
	basic.DataOwner
	CountryTranslationAPI
	CountryID *uuid.UUID `json:"cabin_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:country_translation_unique;not null" swaggertype:"string" format:"uuid"` // Cabin Type id
}

// CountryTranslationAPI struct is writable by request body
type CountryTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:country_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CountryName  *string `json:"country_name,omitempty" gorm:"type:varchar(64)"`
	Nationality  *string `json:"nationality,omitempty" gorm:"type:varchar(64)"`
}
