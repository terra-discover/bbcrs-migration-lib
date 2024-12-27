package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ReservationRatePlan struct {
	basic.Base
	basic.DataOwner
	ReservationRatePlanAPI
}

type ReservationRatePlanAPI struct {
	ReservationID         *uuid.UUID       `json:"reservation_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	RatePlanID            *uuid.UUID       `json:"rate_plan_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	RatePlanCode          *string          `json:"rate_plan_code,omitempty" gorm:"type:varchar(36)"`
	RatePlanName          *string          `json:"rate_plan_name,omitempty" gorm:"type:varchar(64)"`
	RatePlanTypeID        *uuid.UUID       `json:"rate_plan_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	ChargeTypeID          *uuid.UUID       `json:"charge_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	RateTimeUnitID        *uuid.UUID       `json:"rate_time_unit_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	StartDate             *strfmt.DateTime `json:"start_date,omitempty" gorm:"type:timestamptz" format:"date-time"`
	EndDate               *strfmt.DateTime `json:"end_date,omitempty" gorm:"type:timestamptz" format:"date-time"`
	CurrencyID            *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	IsTaxInclusive        *bool            `json:"is_tax_inclusive,omitempty"`
	IsServiceFeeInclusive *bool            `json:"is_service_fee_inclusive,omitempty"`
	IsMealPlanIncluded    *bool            `json:"is_meal_plan_included,omitempty"`
	MealPlanTypeID        *uuid.UUID       `json:"meal_plan_type_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	MealPlanTypeCode      *string          `json:"meal_plan_type_code,omitempty" gorm:"type:varchar(8)"`
	MealPlanTypeName      *string          `json:"meal_plan_type_name,omitempty" gorm:"type:varchar(256)"`
	IsAccrual             *bool            `json:"is_accrual,omitempty"`
	IsRedeemable          *bool            `json:"is_redeemable,omitempty"`
	MarketingLine         *string          `json:"marketing_line,omitempty" gorm:"type:varchar(256)"`
	Comment               *string          `json:"comment,omitempty" gorm:"type:text"`
	Description           *string          `json:"description,omitempty" gorm:"type:text"`
}
