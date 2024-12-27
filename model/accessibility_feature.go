package model

import (
	"github.com/terra-discover/bbcrs-migration-lib/lib"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// AccessibilityFeature Accessibility Feature
type AccessibilityFeature struct {
	basic.Base
	basic.DataOwner
	AccessibilityFeatureAPI
	AccessibilityFeatureTranslation *AccessibilityFeatureTranslation `json:"accessibility_feature_translation,omitempty"`
}

// AccessibilityFeatureAPI Accessibility Feature API
type AccessibilityFeatureAPI struct {
	AccessibilityFeatureCode *int    `json:"accessibility_feature_code,omitempty" gorm:"type:smallint;not null;index:idx_accessibility_feature_code_deleted_at,unique,where:deleted_at is null" example:"1"`                                                    // Accessibility Feature Code
	AccessibilityFeatureName *string `json:"accessibility_feature_name,omitempty" gorm:"type:varchar(256);not null;index:idx_accessibility_feature_code_deleted_at,unique,where:deleted_at is null" example:"Americans with Disabilities Act (ADA) compliance"` // Accessibility Feature Name
}

// Seed Accessibility Feature
func (accessibilityFeature *AccessibilityFeature) Seed() *AccessibilityFeature {
	seed := AccessibilityFeature{
		AccessibilityFeatureAPI: AccessibilityFeatureAPI{
			AccessibilityFeatureCode: lib.Intptr(1),
			AccessibilityFeatureName: lib.Strptr("Americans with Disabilities Act (ADA) compliance"),
		},
	}

	_ = lib.Merge(seed, &accessibilityFeature)
	return accessibilityFeature
}
