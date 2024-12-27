package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RatingTypeTranslation Rating Type Translation
type RatingTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RatingTypeTranslationAPI
	RatingTypeID *uuid.UUID `json:"rating_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:rating_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Rating Type id
}

// RatingTypeTranslationAPI Rating Type Translation API
type RatingTypeTranslationAPI struct {
	LanguageCode   *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:rating_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RatingTypeName *string `json:"rating_type_name,omitempty" gorm:"type:varchar(256)"`                                                             // Rating Type Name
	Provider       *string `json:"provider,omitempty" gorm:"type:varchar(256)"`                                                                     // Provider
	RatingSymbol   *string `json:"rating_symbol,omitempty" gorm:"type:varchar(256)"`                                                                // Rating Symbol
}
