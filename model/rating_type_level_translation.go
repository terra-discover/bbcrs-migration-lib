package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RatingTypeLevelTranslation Rating Type Level Translation
type RatingTypeLevelTranslation struct {
	basic.Base
	basic.DataOwner
	RatingTypeLevelTranslationAPI
	RatingTypeLevelID *uuid.UUID `json:"rating_type_level_id,omitempty" gorm:"type:varchar(36);uniqueIndex:rating_type_level_translation_unique;not null" swaggertype:"string" format:"uuid"` // Rating Type Level id
}

// RatingTypeLevelTranslationAPI Rating Type Level Translation API
type RatingTypeLevelTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:rating_type_level_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RatingTypeLevelName *string `json:"rating_type_level_name,omitempty" gorm:"type:varchar(256)" example:"Low"`                                               // Rating type level name
}
