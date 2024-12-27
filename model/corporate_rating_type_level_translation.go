package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateRatingTypeLevelTranslation Corporate Rating Type Level Translation
type CorporateRatingTypeLevelTranslation struct {
	basic.Base
	basic.DataOwner
	CorporateRatingTypeLevelTranslationAPI
	CorporateRatingTypeLevelID *uuid.UUID `json:"corporate_rating_type_level_id,omitempty" gorm:"type:varchar(36);uniqueIndex:corporate_rating_type_level_translation_unique;not null" swaggertype:"string" format:"uuid"` // Corporate Rating Type Level id
}

// CorporateRatingTypeLevelTranslationAPI Corporate Rating Type Level Translation API
type CorporateRatingTypeLevelTranslationAPI struct {
	LanguageCode                 *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:corporate_rating_type_level_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CorporateRatingTypeLevelName *string `json:"corporate_rating_type_level_name,omitempty" example:"Low" gorm:"type:varchar(256)"`                                               // Corporate Rating Type Level Name
}
