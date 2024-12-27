package model

import basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"

// MealPlanType Meal Plan Type
type MealPlanType struct {
	basic.Base
	basic.DataOwner
	MealPlanTypeAPI
	MealPlanTypeTranslation *MealPlanTypeTranslation `json:"meal_plan_type_translation,omitempty"` // Meal Plan Type Translation
	MealPlanTypeAsset       *MealPlanTypeAsset       `json:"meal_plan_type_asset,omitempty"`       // Meal Plan Type Asset
}

// MealPlanTypeAPI Meal Plan Type API
type MealPlanTypeAPI struct {
	MealPlanTypeCode *int    `json:"meal_plan_type_code,omitempty" gorm:"type:smallint;index:,unique,where:deleted_at is null;not null" example:"123"`             // Meal Plan Type Code
	MealPlanTypeName *string `json:"meal_plan_type_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null" example:"nasi padang"` // Meal Plan Type Name
	// TODO: must review im master-service
	// MealPlanTypeAsset *MealPlanTypeAssetAPI `json:"meal_plan_type_asset,omitempty" gorm:"-"`                                                                                      // Meal Plan Type Asset
}
