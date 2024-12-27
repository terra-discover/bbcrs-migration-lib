package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MealPlanTypeTranslation Meal Plan Type Translation
type MealPlanTypeTranslation struct {
	basic.Base
	basic.DataOwner
	MealPlanTypeTranslationAPI
	MealPlanTypeID *uuid.UUID `json:"meal_plan_type_id,omitempty" gorm:"type:varchar(36);uniqueIndex:meal_plan_type_translation_unique;not null" swaggertype:"string" format:"uuid"` // Meal Plan Type id
}

// MealPlanTypeTranslationAPI Meal Plan Type Translation API
type MealPlanTypeTranslationAPI struct {
	LanguageCode     *string `json:"language_code,omitempty" gorm:"type:varchar(2);uniqueIndex:meal_plan_type_translation_unique;not null" example:"en"` // Language code example: en, id, cn etc...
	MealPlanTypeName *string `json:"meal_plan_type_name,omitempty" gorm:"type:varchar(256)" example:"nasi padang"`                                       // Meal Plan Type Name
}
