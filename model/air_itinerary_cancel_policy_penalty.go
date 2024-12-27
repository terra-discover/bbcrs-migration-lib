package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

type AirItineraryCancelPolicyPenalty struct {
	basic.Base
	basic.DataOwner
	AirItineraryCancelPolicyPenaltyAPI
	AirItineraryCancelPolicy *AirItineraryCancelPolicy `json:"air_itinerary_cancel_policy,omitempty" gorm:"foreignKey:AirItineraryCancelPolicyID;references:ID"`
}

type AirItineraryCancelPolicyPenaltyAPI struct {
	AirItineraryCancelPolicyID *uuid.UUID       `json:"air_itinerary_cancel_policy_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"`
	CancelPolicyID             *uuid.UUID       `json:"cancel_policy_id,omitempty"  gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`
	AbsoluteDeadline           *strfmt.DateTime `json:"absolute_deadline,omitempty" gorm:"type:timestamptz" format:"date-time"`
	OffsetUnitMultiplier       *int             `json:"offset_unit_multiplier,omitempty" gorm:"type:smallint"`
	OffsetDropTimeID           *string          `json:"offset_drop_time_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string"`
	Quantity                   *int             `json:"quantity,omitempty" gorm:"type:smallint"`
	Percent                    *float64         `json:"percent,omitempty"`
	CurrencyID                 *uuid.UUID       `json:"currency_id,omitempty"  gorm:"type:varchar(36)"`
	Amount                     *float64         `json:"amount,omitempty"`
	Description                *string          `json:"description,omitempty" gorm:"type:text"`
}
