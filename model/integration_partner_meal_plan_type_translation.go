package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerMealPlanTypeTranslation Integration Partner Meal Plan Type Translation
type IntegrationPartnerMealPlanTypeTranslation struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerMealPlanTypeTranslationAPI
	IntegrationPartnerMealPlanTypeID *uuid.UUID `json:"integration_partner_meal_plan_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:integration_partner_meal_plan_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Integration Partner Meal Plan Type id
}

// IntegrationPartnerMealPlanTypeTranslationAPI Integration Partner Meal Plan Type Translation API
type IntegrationPartnerMealPlanTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:integration_partner_meal_plan_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MealPlanTypeName *string `json:"meal_plan_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // Meal Plan Type Name
	Description      *string `json:"description,omitempty" gorm:"type:text"`                                                                                                 // Description
	Comment          *string `json:"comment,omitempty" gorm:"type:text"`                                                                                                     // Comment
}
