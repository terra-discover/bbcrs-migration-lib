package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CorporateRatingTypeTranslation Corporate Rating Type Translation
type CorporateRatingTypeTranslation struct {
	basic.Base
	basic.DataOwner
	CorporateRatingTypeTranslationAPI
	CorporateRatingTypeID *uuid.UUID `json:"corporate_rating_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:corporate_rating_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Corporate Rating Type id
}

// CorporateRatingTypeTranslationAPI Corporate Rating Type Translation API
type CorporateRatingTypeTranslationAPI struct {
	LanguageCode            *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:corporate_rating_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	CorporateRatingTypeName *string `json:"corporate_rating_type_name,omitempty" example:"Summary" gorm:"type:varchar(256)"`                                           // Corporate Rating Type Name
}
