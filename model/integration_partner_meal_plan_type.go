package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// IntegrationPartnerMealPlanType Integration Partner Meal Plan Type
type IntegrationPartnerMealPlanType struct {
	basic.Base
	basic.DataOwner
	IntegrationPartnerMealPlanTypeAPI
	IntegrationPartnerMealPlanTypeTranslation *IntegrationPartnerMealPlanTypeTranslation `json:"integration_partner_meal_plan_type_translation,omitempty"` // Integration Partner Meal Plan Type Translation
	MealPlanType                              *MealPlanType                              `json:"meal_plan_type,omitempty"`                                 // Meal Plan Type
}

// IntegrationPartnerMealPlanTypeAPI Integration Partner Meal Plan Type API
type IntegrationPartnerMealPlanTypeAPI struct {
	IntegrationPartnerID *uuid.UUID `json:"integration_partner_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_meal_plan_type__meal_plan_type_code,unique,where:deleted_at is null;not null"` // Integration Partner ID
	MealPlanTypeID       *uuid.UUID `json:"meal_plan_type_id,omitempty" swaggertype:"string" format:"uuid" gorm:"type:varchar(36);index:idx_integration_partner_meal_plan_type__meal_plan_type_code,unique,where:deleted_at is null;not null"`      // Meal Plan Type ID
	MealPlanTypeCode     *string    `json:"meal_plan_type_code,omitempty" gorm:"type:varchar(36);index:idx_integration_partner_meal_plan_type__meal_plan_type_code,unique,where:deleted_at is null;not null"`                                       // Meal Plan Type Code
	MealPlanTypeName     *string    `json:"meal_plan_type_name,omitempty" gorm:"type:varchar(256)"`                                                                                                                                                 // Meal Plan Type Name
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                                                                                                                 // Description
	Comment              *string    `json:"comment,omitempty" gorm:"type:text"`                                                                                                                                                                     // Comment
}

// IntegrationPartnerMealPlanTypeDataGet struct
type IntegrationPartnerMealPlanTypeDataGet struct {
	IntegrationPartnerMealPlanTypeDataGetAPI
	ID                               *uuid.UUID       `json:"id,omitempty"`
	IntegrationPartnerID             *uuid.UUID       `json:"-"`
	IntegrationPartnerMealPlanTypeID *uuid.UUID       `json:"-"`
	MealPlanTypeID                   *uuid.UUID       `json:"-"`
	CreatedAt                        *strfmt.DateTime `json:"created_at,omitempty" swaggertype:"string" format:"date-time"` // created at
	UpdatedAt                        *strfmt.DateTime `json:"updated_at,omitempty" swaggertype:"string" format:"date-time"` // updated at
	Sort                             *int64           `json:"sort,omitempty"`                                               // sort
	Status                           *int64           `json:"status,omitempty"`
}

// IntegrationPartnerMealPlanTypeDataGetAPI struct
type IntegrationPartnerMealPlanTypeDataGetAPI struct {
	MealPlanType                   *MealPlanType                   `json:"meal_plan_type" gorm:"foreignKey:MealPlanTypeID;references:ID"`
	IntegrationPartner             *IntegrationPartner             `json:"integration_partner" gorm:"foreignKey:IntegrationPartnerID;references:ID"`
	IntegrationPartnerMealPlanType *IntegrationPartnerMealPlanType `json:"integration_partner_meal_plan_type" gorm:"foreignKey:IntegrationPartnerMealPlanTypeID;references:ID"`
}
