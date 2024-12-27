package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// RatePlanTypeTranslation Rate Plan Type Translation
type RatePlanTypeTranslation struct {
	basic.Base
	basic.DataOwner
	RatePlanTypeTranslationAPI
	RatePlanTypeID *uuid.UUID `json:"rate_plan_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:rate_plan_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Rate Plan Type id
}

// RatePlanTypeTranslationAPI Rate Plan Type Translation API
type RatePlanTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:rate_plan_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	RatePlanTypeName *string `json:"rate_plan_type_name,omitempty" gorm:"type:varchar(256)" example:"Association"`                                       // Rate Plan Type Name
}
