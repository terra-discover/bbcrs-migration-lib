package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// MealPlanTypeAsset struct
type MealPlanTypeAsset struct {
	basic.Base
	basic.DataOwner
	MealPlanTypeAssetAPI
	MealPlanTypeID        *uuid.UUID             `json:"meal_plan_type_id,omitempty" gorm:"type:varchar(36);index:idx_meal_plan_type_asset__meal_plan_type_id,unique,where:deleted_at is null;not null"` // Meal Plan Type Id
	MultimediaDescription *MultimediaDescription `json:"multimedia_description,omitempty"`                                                                                                               // Multimedia Description
}

// MealPlanTypeAssetAPI struct
type MealPlanTypeAssetAPI struct {
	MultimediaDescriptionID *uuid.UUID `json:"multimedia_description_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Multimedia Description Id
}
