package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AccessibilityFeatureTranslation Accessibility Feature Translation
type AccessibilityFeatureTranslation struct {
	basic.Base
	basic.DataOwner
	AccessibilityFeatureTranslationAPI
	AccessibilityFeatureID *uuid.UUID `json:"accessibility_feature_id,omitempty" gorm:"type:varchar(36);uniqueIndex:accessibility_feature_translation_unique;not null" swaggertype:"string" format:"uuid"` // Accessibility Feature id
}

// AccessibilityFeatureTranslationAPI Accessibility Feature Translation API
type AccessibilityFeatureTranslationAPI struct {
	LanguageCode             *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:accessibility_feature_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	AccessibilityFeatureName *string `json:"accessibility_feature_name,omitempty" gorm:"type:varchar(256)" example:"Disabilities Compliance"`                           // Accessibility Feature Name
}
