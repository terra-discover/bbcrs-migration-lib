package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type ProductPurchaseCancelPolicyPenalty struct {
	basic.Base
	basic.DataOwner
	ProductPurchaseCancelPolicyPenaltyAPI
}

type ProductPurchaseCancelPolicyPenaltyAPI struct {
	ProductPurchaseCancelPolicyID *uuid.UUID       `json:"product_purchase_cancel_policy_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CancelPolicyPenaltyID         *uuid.UUID       `json:"cancel_policy_penalty_id,omitempty" gorm:"type:varchar(36)"`
	AbsoluteDeadline              *strfmt.DateTime `json:"absolute_deadline,omitempty" gorm:"type:timestamptz" format:"date-time"`
	OffsetUnitMultiplier          *int             `json:"offset_unit_multiplier,omitempty" gorm:"type:smallint"`
	OffsetDropTimeID              *string          `json:"offset_drop_time_id,omitempty" gorm:"type:varchar(36)"`
	Quantity                      *int             `json:"quantity,omitempty" gorm:"type:smallint"`
	Percent                       *float64         `json:"percent,omitempty"`
	CurrencyID                    *uuid.UUID       `json:"currency_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" type:"uuid"`
	Amount                        *float64         `json:"amount,omitempty"`
	Description                   *string          `json:"description,omitempty" gorm:"type:text"`
}
