package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AdministrativeCityTranslation Administrative City Translation
type AdministrativeCityTranslation struct {
	basic.Base
	basic.DataOwner
	AdministrativeCityTranslationAPI
	AdministrativeCityID *uuid.UUID `json:"administrative_city_id,omitempty" gorm:"type:varchar(36);uniqueIndex:administrative_city_translation_unique;not null" swaggertype:"string" format:"uuid"` // Administrative City id
}

// AdministrativeCityTranslationAPI Administrative City Translation API
type AdministrativeCityTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:administrative_city_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CityName     *string `json:"city_name,omitempty" gorm:"type:varchar(256)" example:"South Jakarta"`                                                    // City Name
}
