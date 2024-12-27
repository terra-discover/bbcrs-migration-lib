package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// CancelPolicyPenalty Cancel Policy Penalty
type CancelPolicyPenalty struct {
	basic.Base
	basic.DataOwner
	CancelPolicyPenaltyAPI
	CancelPolicyPenaltyTranslation *CancelPolicyPenaltyTranslation `json:"cancel_policy_penalty_translation,omitempty"` // Cancel Policy Penalty Translation
	OffsetDropTime                 *OffsetDropTime                 `json:"offset_drop_time,omitempty"`                  // Offset Drop Time
	CancelPolicy                   *CancelPolicy                   `json:"cancel_policy,omitempty"`                     // Cancel Policy
	Currency                       *Currency                       `json:"currency,omitempty"`                          // Currency
}

// CancelPolicyPenaltyAPI Cancel Policy Penalty API
type CancelPolicyPenaltyAPI struct {
	CancelPolicyID       *uuid.UUID       `json:"cancel_policy_id,omitempty" gorm:"type:varchar(36);not null" swaggertype:"string" format:"uuid"` // Cancel Policy ID
	AbsoluteDeadline     *strfmt.DateTime `json:"absolute_deadline,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`    // Absolute Deadline
	OffsetUnitMultiplier *int             `json:"offset_unit_multiplier,omitempty" gorm:"type:smallint"`                                          // Offset Unit Multiplier
	OffsetDropTimeID     *uuid.UUID       `json:"offset_drop_time_id,omitempty" swaggertype:"string" format:"uuid"`                               // Offset Drop Time ID
	Quantity             *int             `json:"quantity,omitempty" gorm:"type:smallint"`                                                        // Quantity
	Percent              *float64         `json:"percent,omitempty"`                                                                              // Percent
	CurrencyID           *uuid.UUID       `json:"currency_id,omitempty" swaggertype:"string" format:"uuid"`                                       // Currency ID
	Amount               *float64         `json:"amount,omitempty"`                                                                               // Amount
	Description          *string          `json:"description,omitempty" gorm:"type:text"`                                                         // Description
}
