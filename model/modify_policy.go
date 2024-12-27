package model

import (
	"github.com/google/uuid"
	basic "github.com/terra-discover/bbcrs-migration-lib/model/basic"
)

// ModifyPolicy Modify Policy
type ModifyPolicy struct {
	basic.Base
	basic.DataOwner
	ModifyPolicyAPI
	ModifyPolicyTranslation *ModifyPolicyTranslation `json:"modify_policy_translation,omitempty"` // Modify Policy Translation
	OffsetTimeUnit          *OffsetTimeUnit          `json:"offset_time_unit,omitempty"`          // Offset Time Unit
	OffsetDropTime          *OffsetDropTime          `json:"offset_drop_time,omitempty"`          // Offset Drop Time
}

// ModifyPolicyAPI Modify Policy API
type ModifyPolicyAPI struct {
	ModifyPolicyCode     *string    `json:"modify_policy_code,omitempty" gorm:"type:varchar(36);index:,unique,where:deleted_at is null;not null"`  // Modify Policy Code
	ModifyPolicyName     *string    `json:"modify_policy_name,omitempty" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Modify Policy Name
	OffsetTimeUnitID     *uuid.UUID `json:"offset_time_unit_id,omitempty" swaggertype:"string" format:"uuid"`                                      // Offset Time Unit ID
	OffsetUnitMultiplier *int       `json:"offset_unit_multiplier,omitempty" gorm:"type:smallint"`                                                 // Offset Unit Multiplier
	OffsetDropTimeID     *uuid.UUID `json:"offset_drop_time_id,omitempty" swaggertype:"string" format:"uuid"`                                      // Offset Drop Time ID
	NotModifiable        *bool      `json:"not_modifiable,omitempty"`                                                                              // Not Modifiable
	NotRefundable        *bool      `json:"not_refundable,omitempty"`                                                                              // Not Refundable
	Description          *string    `json:"description,omitempty" gorm:"type:text"`                                                                // Description
}
