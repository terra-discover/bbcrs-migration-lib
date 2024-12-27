package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CityTranslation City Translation
type CityTranslation struct {
	basic.Base
	basic.DataOwner
	CityTranslationAPI
	CityID *uuid.UUID `json:"city_id,omitempty" gorm:"type:varchar(36);uniqueIndex:city_translation_unique;not null" swaggertype:"string" format:"uuid"` // City id
}

// CityTranslationAPI City Translation API
type CityTranslationAPI struct {
	LanguageCode *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:city_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CityName     *string `json:"city_name,omitempty" gorm:"type:varchar(256)" example:"Jakarta"`                                           // City Name
}
