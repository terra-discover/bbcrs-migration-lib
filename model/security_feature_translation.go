package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// SecurityFeatureTranslation Security Feature Translation
type SecurityFeatureTranslation struct {
	basic.Base
	basic.DataOwner
	SecurityFeatureTranslationAPI
	SecurityFeatureID *uuid.UUID `json:"security_feature_id,omitempty" gorm:"type:varchar(36);uniqueIndex:security_feature_translation_unique;not null" swaggertype:"string" format:"uuid"` // Security Feature id
}

// SecurityFeatureTranslationAPI Security Feature Translation API
type SecurityFeatureTranslationAPI struct {
	LanguageCode        *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:security_feature_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	SecurityFeatureName *string `json:"security_feature_name,omitempty" gorm:"type:varchar(256)" example:"2nd lock on guest doors"`                           // Security Feature Name
}
